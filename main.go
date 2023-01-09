package main

import (
	"fmt"
	"log"

	"github.com/NaokiYazawa/go-simple-server/server"
)

const (
	host = "0.0.0.0"
	port = "5000"
)

func main() {
	s := server.New()
	err := s.Start(fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		log.Fatalf("server stopped with error: %s", err)
	}

	log.Printf("server stopped successfully")
}
