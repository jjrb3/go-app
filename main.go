package main

import (
	"log"

	"github.com/jjrb3/go-app/db"
	"github.com/jjrb3/go-app/handlers"
)

func main() {
	if db.CheckConnection() == 0 {
		log.Fatal("Without connection to DB")
		return
	}

	handlers.Handlers()
}
