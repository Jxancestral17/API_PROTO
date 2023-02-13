package main

import (
	"log"

	"github.com/Jxancestral17/API_PROTO/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
