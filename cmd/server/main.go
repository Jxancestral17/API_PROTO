package main

import (
	"log"

	"github.com/Jxancestral17/distributed_services_learn/distributed_services/internal/server"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
