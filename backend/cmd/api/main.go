package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	database "github.com/Duuendy/HoraMarcada/backend/internal/database"
	dbConfig "github.com/Duuendy/HoraMarcada/backend/internal/database"
	apphttp "github.com/Duuendy/HoraMarcada/backend/internal/http"

	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func main() {
	//DB concection
	db, err = sql.Open("postgres", dbConfig.DataSourceName)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Printf("PING\n")

	repo := &database.ServiceRepo{DB: db}

	mux := apphttp.Router(repo)

	addr := ":8080"
	if p := strings.TrimSpace(os.Getenv("PORT")); p != "" {
		if strings.HasPrefix(p, ":") {
			addr = p
		} else {
			addr = ":" + p
		}
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	serveErr := make(chan error, 10)
	go func() {
		fmt.Printf("Starting server on %s (set PORT to avoid clashing with another service on 8080)\n", addr)
		serveErr <- http.ListenAndServe(addr, mux)
	}()

	select {
	case sig := <-sigs:
		fmt.Printf("PONG - Encerrando (%s)\n", sig.String())
	case err := <-serveErr:
		// se o servidor falhar, você também encerra
		fmt.Printf("PONG - Servidor caiu: %v\n", err)
	}
}
