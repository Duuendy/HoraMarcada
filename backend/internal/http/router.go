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
	// Instanciamos o handler passando o banco
    serviceHandler := &ServiceHandler{DB: db}

    // Registramos o método diretamente
    mux.HandleFunc("/services/list", serviceHandler.List)
	return mux	
}
	// mux.HandleFunc("/services/", GetServiceHandler)
