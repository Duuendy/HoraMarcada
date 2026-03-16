package http

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	dto "github.com/Duuendy/HoraMarcada/backend/internal/http/dto/service"
	resp "github.com/Duuendy/HoraMarcada/backend/internal/http/response"
	service "github.com/Duuendy/HoraMarcada/backend/internal/service"
)

type ServiceHandler struct {
	DB *sql.DB
}

func CreateServiceHandler(h http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		resp.ResponseError(h, http.StatusMethodNotAllowed, &resp.APIError{
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
		})
		return
	}
	req := dto.DTOCreateServiceRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		resp.ResponseError(h, http.StatusBadRequest, &resp.APIError{
			Code:    http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}
	name := strings.TrimSpace(req.Name)
	if name == "" {
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
	id := service.CreateService(
		req.Name,
		req.PriceCent,
		req.TimeMinutes,
		req.IsMaintenance,
	)
	responseService := dto.DTOCreateServiceResponse{
		Id:            id,
		Name:          req.Name,
		PriceCent:     req.PriceCent,
		TimeMinutes:   req.TimeMinutes,
		IsMaintenance: req.IsMaintenance,
	}
	resp.ResponseSuccess(h, responseService)
}

func (sh *ServiceHandler) List(h http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		resp.ResponseError(h, http.StatusMethodNotAllowed, &resp.APIError{
			Code:    http.StatusMethodNotAllowed,
			Message: "Method not allowed",
		})
		return
	}

	servItens, err := service.ListServices(sh.DB)
	if err != nil {
		resp.ResponseError(h, http.StatusInternalServerError, &resp.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Database error!!",
		})
		return
	}
	resp.ResponseSuccess(h, servItens)
}

// func GetServiceHandler(h http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodGet {
// 		resp.ResponseError(h, http.StatusMethodNotAllowed, &resp.APIError{
// 			Code:    http.StatusMethodNotAllowed,
// 			Message: "Method not allowes",
// 		})
// 		return
// 	}
// 	path := strings.TrimPrefix(r.URL.Path, "/services/")
// 	id, err := strconv.Atoi(path)
// 	if err != nil {
// 		resp.ResponseError(h, http.StatusBadRequest, &resp.APIError{
// 			Code:    http.StatusBadRequest,
// 			Message: "Invalid request payload",
// 		})
// 		return
// 	}

// 	serviceModel, servFind := service.GetServiceByID(id)
// 	if !servFind {
// 		resp.ResponseError(h, http.StatusNotFound, &resp.APIError{
// 			Code:    http.StatusNotFound,
// 			Message: "Service not found",
// 		})
// 		return
// 	}

// 	response := mapper.ToServiceItem(serviceModel)
// 	resp.ResponseSuccess(h, response)
// }
