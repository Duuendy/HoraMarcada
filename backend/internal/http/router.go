package http

import (
	"net/http"
)

// Router cria e retorna o roteador da aplicação
func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Handler)
	mux.HandleFunc("/services", CreateServiceHandler)
	mux.HandleFunc("/services/list", ListServicesHandler)
	mux.HandleFunc("/services/", GetServiceHandler)
	return mux
}