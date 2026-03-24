package http

import (
	"encoding/json"
	"net/http"
	"strings"

	database "github.com/Duuendy/HoraMarcada/backend/internal/database"
	"github.com/Duuendy/HoraMarcada/backend/internal/domain"
	dto "github.com/Duuendy/HoraMarcada/backend/internal/http/dto/dto_service"
	"github.com/Duuendy/HoraMarcada/backend/internal/http/mapper"
	resp "github.com/Duuendy/HoraMarcada/backend/internal/http/response"
)

type ServiceHandler struct {
	Repository database.ServiceRepository
}

func (sh *ServiceHandler) Create(h http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if r.Method != http.MethodPost {
		resp.ResponseError(h, http.StatusMethodNotAllowed, &resp.APIError{
			Code:    http.StatusMethodNotAllowed,
			Message: "Use POST with JSON body (opening this URL in the browser sends GET only)",
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
	id, err := sh.Repository.Create(domain.ServiceModel{
		Name:          req.Name,
		PriceCent:     req.PriceCent,
		TimeMinutes:   req.TimeMinutes,
		IsMaintenance: req.IsMaintenance,
	})
	if err != nil {
		resp.ResponseError(h, http.StatusInternalServerError, &resp.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Erro ao gravar serviço",
		})
		return
	}
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

	servItens, err := sh.Repository.List()

	if err != nil {
		resp.ResponseError(h, http.StatusInternalServerError, &resp.APIError{
			Code:    http.StatusInternalServerError,
			Message: "Erro ao buscar dados no banco",
		})
		return
	}
	items := make([]dto.ServiceItem, 0, len(servItens))
	for _, s := range servItens {
		items = append(items, mapper.ToServiceItem(s))
	}
	resp.ResponseSuccess(h, dto.DTOListServiceResponse{Items: items})
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
