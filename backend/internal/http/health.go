package http

import (
	"fmt"
	"net/http"
)
// Handler responde ao endpoint /health
func Handler(h http.ResponseWriter, r *http.Request) {
	fmt.Fprint(h, "Hello Word!!!")
}