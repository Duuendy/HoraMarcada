package main

import (
	"fmt"
	"net/http"

	apphttp "github.com/Duuendy/HoraMarcada/backend/internal/http"
)

func main() {
	fmt.Printf("Starting server on :8080\n")

	mux := apphttp.Router()

	http.ListenAndServe(":8080", mux)
}