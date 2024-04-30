package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		requestID := request.Header.Get("X-Request-ID")

		log.WithFields(log.Fields{
			"X-Request-ID": requestID,
			"method":       request.Method,
			"request":      request.RequestURI,
			"hostname":     request.Host,
		}).Info("Started handling request.")

		next(response, request, params)
	}
}
