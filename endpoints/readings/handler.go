package readings

import (
	"encoding/json"
	"fmt"
	"io"
	"joi-energy-golang/api"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"joi-energy-golang/domain"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) StoreReadings(response http.ResponseWriter, request *http.Request, urlParams httprouter.Params) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		api.Error(response, request, fmt.Errorf("read request body failed: %w", err), http.StatusBadRequest)
		return
	}

	var readings domain.StoreReadings

	if err := json.Unmarshal(body, &readings); err != nil {
		api.Error(response, request, fmt.Errorf("unmarshal request body failed: %w", err), http.StatusBadRequest)
		return
	}

	err = readings.Validate()
	if err != nil {
		api.Error(response, request, err, http.StatusBadRequest)
		return
	}

	h.service.StoreReadings(readings.SmartMeterId, readings.ElectricityReadings)

	api.Success(response, request, nil)
}

func (h *Handler) GetReadings(response http.ResponseWriter, request *http.Request, urlParams httprouter.Params) {
	smartMeterId := urlParams.ByName("smartMeterId")

	err := validateSmartMeterId(smartMeterId)
	if err != nil {
		api.Error(response, request, err, http.StatusBadRequest)
		return
	}

	readings := h.service.GetReadings(smartMeterId)

	result := domain.StoreReadings{
		SmartMeterId:        smartMeterId,
		ElectricityReadings: readings,
	}

	api.SuccessJson(response, request, result)
}
