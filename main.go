package main

import (
	"api/internal/todo"
	"api/internal/transport"
	"log"
)

func main() {

	svc := todo.NewService()
	server := transport.NewServer(svc)

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}

}
