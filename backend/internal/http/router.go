package http

import (
	"database/sql"
	"net/http"
)

// Router cria e retorna o roteador da aplicação
func Router(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Handler)
	mux.HandleFunc("/services", CreateServiceHandler)
	mux.HandleFunc("/services/list", ListServicesHandler)
	// mux.HandleFunc("/services/list", func(w http.ResponseWriter, r *http.Request) {
	// 	ListServicesHandler(w,r,db)
	// })
	mux.HandleFunc("/services/", GetServiceHandler)

	return mux
}