package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	dbConfig "github.com/Duuendy/HoraMarcada/backend/db"
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

	var dbName string
	if err := db.QueryRow("select current_database()").Scan(&dbName); err != nil {
		panic(err)
	}
	fmt.Println("DB atual:", dbName)

	var schema string
	if err := db.QueryRow("select current_schema()").Scan(&schema); err != nil {
		panic(err)
	}
	fmt.Println("Schema atual:", schema)

	var addr string
	if err := db.QueryRow("select inet_server_addr()::text || ':' || inet_server_port()::text").Scan(&addr); err != nil {
		panic(err)
	}
	fmt.Println("Servidor Postgres:", addr)

	fmt.Printf("PING\n")

	mux := apphttp.Router(db)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	serveErr := make(chan error, 10)
	go func() {
		fmt.Printf("Starting server on :8080\n")
		serveErr <- http.ListenAndServe(":8080", mux)
	}()

	select {
	case sig := <-sigs:
		fmt.Printf("PONG - Encerrando (%s)\n", sig.String())
	case err := <-serveErr:
		// se o servidor falhar, você também encerra
		fmt.Printf("PONG - Servidor caiu: %v\n", err)
	}
}
