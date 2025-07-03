package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error  string `json:"error"`
	Status int    `json:"status"`
}

type SuccessResponse struct {
	Data   string `json:"data"`
	Status int    `json:"status"`
}

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func GeneralError(err error, status int) ErrorResponse {

	return ErrorResponse{
		Status: status,
		Error:  err.Error(),
	}
}

func CustomResponse(msg string, status int) SuccessResponse {
	return SuccessResponse{
		Data:   msg,
		Status: status,
	}
}
