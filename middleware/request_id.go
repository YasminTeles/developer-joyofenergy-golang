package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func RequestIDMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(response http.ResponseWriter, request *http.Request, params httprouter.Params) {
		if request.Header.Get("X-Request-ID") == "" {
			request.Header.Set("X-Request-ID", uuid.New().String())
		}

		response.Header().Set("X-Request-ID", request.Header.Get("X-Request-ID"))

		next(response, request, params)
	}
}
