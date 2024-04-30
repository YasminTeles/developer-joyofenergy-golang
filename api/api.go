package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"joi-energy-golang/domain"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func SuccessJson(w http.ResponseWriter, r *http.Request, data interface{}) {
	jsonMsg, err := json.Marshal(data)
	if err != nil {
		Error(w, r, fmt.Errorf("serialising response failed: %w", err), 500)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		Success(w, r, jsonMsg)
	}
}

func Success(w http.ResponseWriter, r *http.Request, jsonMsg []byte) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if _, err := w.Write(jsonMsg); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error writing response.")
	}

	log.WithFields(log.Fields{
		"method":   r.Method,
		"request":  r.RequestURI,
		"port":     r.RemoteAddr,
		"httpCode": "200",
	}).Info("Success writing response.")

}

func Error(w http.ResponseWriter, r *http.Request, err error, code int) {
	if code == 0 {
		code = toHTTPStatusCode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)

	if err == nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Nil Error")
	}

	logErr := err
	errorMsgJSON, err := json.Marshal(domain.ErrorResponse{Message: err.Error()})
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error Marshal.")

	} else {
		if _, err = w.Write(errorMsgJSON); err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Error("Error writing response.")
		}
	}

	log.WithFields(log.Fields{
		"method":   r.Method,
		"request":  r.RequestURI,
		"port":     r.RemoteAddr,
		"httpCode": code,
		"error":    logErr.Error(),
	}).Error("Error writing response.")
}

func toHTTPStatusCode(err error) int {
	switch {
	case errors.Is(err, domain.ErrMissingArgument):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrInvalidMessageType):
		return http.StatusBadRequest
	case errors.Is(err, domain.ErrNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
