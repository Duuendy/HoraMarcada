package response

import (
	"encoding/json"
	"net/http"
)
type APIResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   *APIError   `json:"error,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Meta struct {
	TotalCount int `json:"total_count,omitempty"`
	Page       int `json:"page,omitempty"`
	PageSize   int `json:"page_size,omitempty"`
}

func ResponseSuccess(h http.ResponseWriter, data interface{}) {
	h.Header().Set("Content-Type", "application/json; charset=utf-8")
	h.WriteHeader(http.StatusOK)
	response := APIResponse{
		Success: true,
		Data: data,
		Meta: nil,
	}
	json.NewEncoder(h).Encode(response)
}

func ResponseError(h http.ResponseWriter, code int, apiErr *APIError) {
	h.Header().Set("Content-Type", "application/json; charset=utf-8")
	h.WriteHeader(code)
	response := APIResponse{
		Success: false,
		Error: apiErr,
		Meta: nil,
	}
	json.NewEncoder(h).Encode(response)
}