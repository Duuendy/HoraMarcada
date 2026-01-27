package http

import (
	"encoding/json"
	"net/http"
	"strings"

	dto "github.com/Duuendy/HoraMarcada/backend/internal/http/dto/service"
	resp "github.com/Duuendy/HoraMarcada/backend/internal/http/response"
)

func CreateServiceHandler(h http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost{
		resp.ResponseError(h, http.StatusMethodNotAllowed, &resp.APIError{
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
		})
		return
	}
	req := dto.CreateServiceRequest{
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
	responseService := dto.CreateServiceResponse{	
		Id: 1,
		Name: name,
		PriceCent: req.PriceCent,
		TimeMinutes: req.TimeMinutes,
		IsMaintenance: req.IsMaintenance,
	}
	resp.ResponseSuccess(h, responseService)
}