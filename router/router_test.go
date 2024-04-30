package router

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"joi-energy-golang/domain"
)

// nolint: paralleltest
func TestServer(t *testing.T) {
	t.Setenv("PORT", "8081")

	server := NewServer()
	go func() {
		err := server.ListenAndServe()
		if err != http.ErrServerClosed {
			t.Error(err)
		}
	}()
	defer server.Close()

	// Wait 50 milliseconds for server to start listening to requests
	time.Sleep(50 * time.Millisecond)

	resp, err := http.Get("http://localhost:8081/readings/read/smartMeterId")

	assert.NoError(t, err)
	defer resp.Body.Close()

	expectedContentType := "application/json"
	actualContentType := resp.Header.Get("Content-Type")
	assert.Equalf(t, expectedContentType, actualContentType, "handler returned wrong content-type: got %v want %v", actualContentType, expectedContentType)
}

func TestEndpointEndpointSuccess(t *testing.T) {
	t.Parallel()

	testHandler := newHandler()

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest(http.MethodGet, "/readings/read/smartMeterId", nil)
	req.Header.Add("Content-Type", "application/json")

	testHandler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	expected := domain.StoreReadings{
		SmartMeterId:        "smartMeterId",
		ElectricityReadings: nil,
	}

	var actual domain.StoreReadings
	err := json.Unmarshal(rr.Body.Bytes(), &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestHealthcheckEndPoint(t *testing.T) {
	t.Parallel()

	testHandler := newHandler()

	rr := httptest.NewRecorder()

	request, _ := http.NewRequest(http.MethodGet, "/healthcheck", nil)

	testHandler.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)

	byteBody, _ := io.ReadAll(rr.Body)

	message := strings.Trim(string(byteBody), "\n")

	assert.Equal(t, "Working!", message)
}

// nolint: paralleltest
func TestGetDefaultPort(t *testing.T) {
	t.Setenv("PORT", "")

	actual := getListeningPort()

	assert.Equal(t, ":8080", actual)
}

// nolint: paralleltest
func TestGetCustomPort(t *testing.T) {
	t.Setenv("PORT", "4000")

	actual := getListeningPort()

	assert.Equal(t, ":4000", actual)
}

func TestElectricityCostEndPoint(t *testing.T) {
	t.Parallel()

	testHandler := newHandler()
	rr := httptest.NewRecorder()

	body := `{"smartMeterId": "smart-meter-0", "electricityReadings": []}`

	request, _ := http.NewRequest(http.MethodPost, "/price-plans/estimate", strings.NewReader(body))

	testHandler.ServeHTTP(rr, request)
	assert.Equal(t, http.StatusOK, rr.Code)
}
