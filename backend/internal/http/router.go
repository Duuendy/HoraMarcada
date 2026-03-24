package http

import (
	"net/http"

	"github.com/Duuendy/HoraMarcada/backend/internal/database"
)

// Router cria e retorna o roteador da aplicação
func Router(repo database.ServiceRepository) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", Handler)
	// Instanciamos o handler passando o banco
	serviceHandler := &ServiceHandler{Repository: repo}
	// Registramos o método diretamente
	mux.HandleFunc("/services/list", serviceHandler.List)
	mux.HandleFunc("/services/create", serviceHandler.Create)
	return mux
}

// mux.HandleFunc("/services/", GetServiceHandler)
