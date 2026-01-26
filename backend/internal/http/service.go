package http

import (
	"encoding/json"
	"net/http"
	"strings"

	resp "github.com/Duuendy/HoraMarcada/backend/internal/http/response"
)
type CreateServiceRequest struct {
	Name		 	string	`json:"name"`
	PriceCent		int		`json:"price_cent"`
	TimeMinutes    	int		`json:"time_minutes"` // in minutes
	IsMaintenance	bool	`json:"is_maintenance"`
}

func CreateServiceHandler(h http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost{
		resp.ResponseError(h, http.StatusMethodNotAllowed, &resp.APIError{
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
		})
		return
	}
	req := CreateServiceRequest{
	}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {		
		resp.ResponseError(h, http.StatusBadRequest, &resp.APIError{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}
	name := strings.TrimSpace(req.Name)
	if name == ""  {
		resp.ResponseError(h, http.StatusBadRequest, &resp.APIError{
			Code:    http.StatusBadRequest,
			Message: "Name is required",
		})
		return
	}
	if req.PriceCent <= 0 || req.TimeMinutes <= 0 {
		resp.ResponseError(h, http.StatusBadRequest, &resp.APIError{
			Code:    http.StatusBadRequest,
			Message: "Price and Time must be greater than zero",
		})
		return
	}
	resp.ResponseSuccess(h, map[string]string{
		"message": "Service created successfully",
	})	
}