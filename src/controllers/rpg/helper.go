package rpg

import (
	"encoding/json"
	"net/http"
)

func handleError(err error, message string, response http.ResponseWriter) http.ResponseWriter {
	errorBody, _ := json.Marshal(responseError{Error: message})
	response.WriteHeader(http.StatusBadRequest)
	response.Write(errorBody)
	return response
}
