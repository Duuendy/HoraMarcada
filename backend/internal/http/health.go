package http

import (
	"net/http"

	resp "github.com/Duuendy/HoraMarcada/backend/internal/http/response"
)

func Handler(h http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("fail") == "true" {
		resp.ResponseError(h, http.StatusServiceUnavailable, &resp.APIError{
			Code:    http.StatusServiceUnavailable,
			Message: "Service Low",
		})
		return
	}
	resp.ResponseSuccess(h, map[string]string{
		"status": "OK",
	})
}
