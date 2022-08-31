package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jjrb3/go-app/middlewares"
	"github.com/rs/cors"
)

// Handlers is a function that set port and listen the server.
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "3000"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
