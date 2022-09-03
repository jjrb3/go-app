package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jjrb3/go-app/db"
	"github.com/jjrb3/go-app/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error in the data received"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "The user email is required", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "The password would be minimum 6 characters", http.StatusBadRequest)
		return
	}

	_, exist, _ := db.FindByEmail(t.Email)

	if exist == true {
		http.Error(w, "Already the user is registered", http.StatusBadRequest)
		return
	}

	_, status, err := db.Insert(t)

	if err != nil {
		http.Error(w, "Error when registering the user "+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "User could not be registered"+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
