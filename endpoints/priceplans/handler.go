package priceplans

import (
	"encoding/json"
	"fmt"
	"io"
	"joi-energy-golang/api"
	"joi-energy-golang/domain"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CompareAll(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	smartMeterId := urlParams.ByName("smartMeterId")
	err := validateSmartMeterId(smartMeterId)
	if err != nil {
		api.Error(w, r, err, http.StatusBadRequest)
		return
	}
	result, err := h.service.CompareAllPricePlans(smartMeterId)
	if err != nil {
		api.Error(w, r, err, 0)
		return
	}
	api.SuccessJson(w, r, result)
}

func (h *Handler) Recommend(w http.ResponseWriter, r *http.Request, urlParams httprouter.Params) {
	smartMeterId := urlParams.ByName("smartMeterId")
	err := validateSmartMeterId(smartMeterId)
	if err != nil {
		api.Error(w, r, err, http.StatusBadRequest)
		return
	}
	limitString := r.URL.Query().Get("limit")
	limit, err := strconv.ParseUint(limitString, 10, 64)
	if limitString != "" && err != nil {
		api.Error(w, r, err, http.StatusBadRequest)
		return
	}
	result, err := h.service.RecommendPricePlans(smartMeterId, limit)
	if err != nil {
		api.Error(w, r, err, 0)
		return
	}
	api.SuccessJson(w, r, result)
}

func (h *Handler) ElectricityCost(response http.ResponseWriter, request *http.Request, urlParams httprouter.Params) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		api.Error(
			response,
			request,
			fmt.Errorf("Read request body failed: %w", err),
			http.StatusBadRequest)
		return
	}

	var readings domain.StoreReadings
	if err := json.Unmarshal(body, &readings); err != nil {
		api.Error(
			response,
			request,
			fmt.Errorf("Unmarshal request body failed: %w", err),
			http.StatusBadRequest)
		return
	}

	err = validateSmartMeterId(readings.SmartMeterId)
	if err != nil {
		api.Error(response, request, err, http.StatusBadRequest)
		return
	}

	allPricePlans, err := h.service.CompareAllPricePlans(readings.SmartMeterId)
	if err != nil {
		api.Error(response, request, err, 0)
		return
	}

	recommendations, err := h.service.RecommendPricePlans(readings.SmartMeterId, 0)
	if err != nil {
		api.Error(response, request, err, 0)
		return
	}

	electricityCost := getElectricityCost(readings, allPricePlans)

	api.SuccessJson(response, request, &domain.ElectricityCost{
		ElectricityCost: electricityCost,
		Recommendations: recommendations.Recommendations,
	})
}

func getElectricityCost(readings domain.StoreReadings, allPricePlans domain.PricePlanComparisons) float64 {
	electricityAmount := sumElectricity(readings)
	return allPricePlans.PricePlanComparisons[allPricePlans.PricePlanId] * electricityAmount
}

func sumElectricity(readings domain.StoreReadings) float64 {
	electricityAmount := 0.0

	for _, electricityReadings := range readings.ElectricityReadings {
		electricityAmount += electricityReadings.Reading
	}

	return electricityAmount
}
