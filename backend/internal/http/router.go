package http

import (
	"net/http"
)

// Router cria e retorna o roteador da aplicação
func Router() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Handler)
	mux.HandleFunc("/service", CreateServiceHandler)
	return mux
}