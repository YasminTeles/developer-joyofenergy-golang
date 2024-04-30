package priceplans

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)

func callEndpoint(handler http.HandlerFunc, url string, t *testing.T) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatalf("request creation failed: %s", err)
	}
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
	return rr
}

func TestCompareAllPricePlansReturnResultFromService(t *testing.T) {
	t.Parallel()

	s := &MockService{}
	h := NewHandler(s)
	params := httprouter.Params{{Key: "smartMeterId", Value: "123"}}
	compareAllPricePlans := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		h.CompareAll(writer, request, params)
	})

	rr := callEndpoint(compareAllPricePlans, "/price-plans/compare-all/123", t)
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned status code %v on valid request", rr.Code)
	var data domain.PricePlanComparisons

	err := json.Unmarshal(rr.Body.Bytes(), &data)
	assert.NoError(t, err)

	assert.Equal(t, domain.PricePlanComparisons{}, data)
}

func TestCompareAllPricePlansHandleServiceError(t *testing.T) {
	t.Parallel()

	s := &MockService{err: errors.New("oops")}
	h := NewHandler(s)
	params := httprouter.Params{{Key: "smartMeterId", Value: "123"}}
	compareAllPricePlans := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		h.CompareAll(writer, request, params)
	})

	rr := callEndpoint(compareAllPricePlans, "/price-plans/compare-all/123", t)
	assert.NotEqual(t, http.StatusOK, rr.Code, "handler returned status code %v on failing request", rr.Code)

	var response domain.ErrorResponse

	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, domain.ErrorResponse{Message: "oops"}, response)
}

// func TestElectricityCostHandler(t *testing.T) {
// 	t.Parallel()

// 	s := &MockService{}
// 	h := NewHandler(s)
// 	params := httprouter.Params{{Key: "smartMeterId", Value: ""}}
// 	recommendPricePlans := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
// 		h.Recommend(writer, request, params)
// 	})

// 	rr := callEndpoint(recommendPricePlans, "/price-plans/estimate/", t)
// 	assert.NotEqual(t, http.StatusOK, rr.Code, "handler returned status code %v on failing request", rr.Code)

// 	var response domain.ErrorResponse

// 	err := json.Unmarshal(rr.Body.Bytes(), &response)
// 	assert.NoError(t, err)

// 	assert.Equal(t, domain.ErrorResponse{Message: "cannot be blank"}, response)
// }

func TestCompareAllPricePlansHandlerWithInvalidInput(t *testing.T) {
	t.Parallel()

	s := &MockService{}
	h := NewHandler(s)
	params := httprouter.Params{{Key: "smartMeterId", Value: ""}}
	recommendPricePlans := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		h.Recommend(writer, request, params)
	})

	rr := callEndpoint(recommendPricePlans, "/price-plans/recommend/", t)
	assert.NotEqual(t, http.StatusOK, rr.Code, "handler returned status code %v on failing request", rr.Code)

	var response domain.ErrorResponse

	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, domain.ErrorResponse{Message: "cannot be blank"}, response)
}

func TestSumElectricity(t *testing.T) {
	t.Parallel()

	readings := domain.StoreReadings{
		SmartMeterId: "smartMeterId",
		ElectricityReadings: []domain.ElectricityReading{
			{
				Time:    time.Now(),
				Reading: 2.449985365357211,
			},
			{
				Time:    time.Now(),
				Reading: 2.435749415846827,
			},
		},
	}

	amount := sumElectricity(readings)

	assert.Equal(t, amount, 4.885734781204038)
}

func TestGetElectricityCost(t *testing.T) {
	t.Parallel()

	readings := domain.StoreReadings{
		SmartMeterId: "smartMeterId",
		ElectricityReadings: []domain.ElectricityReading{
			{
				Time:    time.Now(),
				Reading: 2.449985365357211,
			},
			{
				Time:    time.Now(),
				Reading: 2.435749415846827,
			},
		},
	}

	allPricePlans := domain.PricePlanComparisons{
		PricePlanId: "price-plan-0",
		PricePlanComparisons: map[string]float64{
			"price-plan-0": 210.038654450053,
			"price-plan-1": 42.0077308900106,
			"price-plan-2": 21.0038654450053,
		},
	}

	cost := getElectricityCost(readings, allPricePlans)

	assert.Equal(t, cost, 1026.1931594439202)
}

type MockService struct {
	err error
	Service
}

func (s *MockService) CompareAllPricePlans(smartMeterId string) (domain.PricePlanComparisons, error) {
	return domain.PricePlanComparisons{}, s.err
}
