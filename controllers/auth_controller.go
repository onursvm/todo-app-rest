package controllers

import (
	"encoding/json"
	"net/http"
	"todo-app/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	json.NewDecoder(r.Body).Decode(&creds)

	for _, user := range utils.Data.Users {
		if user.Username == creds.Username && user.Password == creds.Password {
			token, _ := utils.GenerateJWT(user.Username, user.Role)
			json.NewEncoder(w).Encode(map[string]string{"token": token})
			return
		}
	}
	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}
