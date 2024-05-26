package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	common "github.com/joshua468/commons"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8000")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server")
	}
}
