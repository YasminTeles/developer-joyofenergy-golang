package middleware

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func LoggingMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		log.WithFields(log.Fields{
			"method":   request.Method,
			"request":  request.RequestURI,
			"hostname": request.Host,
		}).Info("Started handling request.")

		next(response, request, params)
	}
}
