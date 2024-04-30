package router

import (
	"fmt"
	"joi-energy-golang/api"
	"joi-energy-golang/endpoints/priceplans"
	"joi-energy-golang/endpoints/readings"
	"joi-energy-golang/endpoints/standard"
	"joi-energy-golang/middleware"
	"joi-energy-golang/repository"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func NewServer() *http.Server {
	return &http.Server{
		Addr:         getListeningPort(),
		Handler:      newHandler(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}

func getListeningPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.WithFields(log.Fields{
			"port": port,
		}).Info("Defaulting to port...")

	}

	return fmt.Sprintf(":%s", port)
}

func addRoutes(router *httprouter.Router) {
	accounts := repository.NewAccounts(defaultSmartMeterToPricePlanAccounts())
	meterReadings := repository.NewMeterReadings(
		defaultMeterElectricityReadings(),
	)
	pricePlans := repository.NewPricePlans(
		defaultPricePlans(),
		&meterReadings,
	)

	readingsHandler := readings.NewHandler(&meterReadings)
	pricePlanHandler := priceplans.NewHandler(priceplans.NewService(&pricePlans, &accounts))

	router.GET("/healthcheck", middleware.RequestIDMiddleware(middleware.LoggingMiddleware(standard.Healthcheck)))
	router.GET("/version", middleware.RequestIDMiddleware(middleware.LoggingMiddleware(standard.Version)))

	router.POST("/readings/store", middleware.RequestIDMiddleware(middleware.LoggingMiddleware(readingsHandler.StoreReadings)))
	router.GET("/readings/read/:smartMeterId", middleware.RequestIDMiddleware(middleware.LoggingMiddleware(readingsHandler.GetReadings)))

	router.GET("/price-plans/compare-all/:smartMeterId", middleware.RequestIDMiddleware(middleware.LoggingMiddleware(pricePlanHandler.CompareAll)))
	router.GET("/price-plans/recommend/:smartMeterId", middleware.RequestIDMiddleware(middleware.LoggingMiddleware(pricePlanHandler.Recommend)))
	router.POST("/price-plans/estimate", middleware.RequestIDMiddleware(middleware.LoggingMiddleware(pricePlanHandler.ElectricityCost)))
}

func newHandler() http.Handler {
	r := httprouter.New()
	addRoutes(r)

	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		}
		w.WriteHeader(http.StatusNoContent)
	})

	r.PanicHandler = func(w http.ResponseWriter, r *http.Request, err interface{}) {
		log.WithFields(log.Fields{
			"error": err,
		}).Panic("whoops! My handler has run into a panic.")

		api.Error(w, r, fmt.Errorf("whoops! My handler has run into a panic"), http.StatusInternalServerError)
	}

	r.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.Error(w, r, fmt.Errorf("we have OPTIONS for youm but %v is not among them", r.Method), http.StatusMethodNotAllowed)
	})

	r.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept"), "text/html") {
			sendBrowserDoc(w, r)
			return
		}
		api.Error(w, r, fmt.Errorf("whatever route you've been looking for, it's not here"), http.StatusNotFound)
	})

	return r
}

func sendBrowserDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusUnsupportedMediaType)
	b, err := os.ReadFile("browser.htm")
	if err != nil {
		api.Error(w, r, fmt.Errorf("read browser.htm failed: %w", err), http.StatusInternalServerError)
	}
	_, err = w.Write(b)
	if err != nil {
		api.Error(w, r, fmt.Errorf("send browser.htm failed: %w", err), http.StatusInternalServerError)
	}
}
