package middlewares

import (
	"net/http"

	"github.com/jjrb3/go-app/db"
)

// CheckDB is the middleware that check the database status.
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if db.CheckConnection() == 0 {
			http.Error(rw, "Connect lost with the Data base", 500)
			return
		}

		next.ServeHTTP(rw, r)
	}
}
