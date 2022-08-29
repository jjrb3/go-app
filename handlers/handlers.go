package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Handlers is a function that set port and listen the server.
func Handlers() {
	router := mux.NewRouter()

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "3000"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
