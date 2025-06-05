package handlers

import (
	"encoding/json"
	"net/http"

	"myproject/internal/service"
	"myproject/models"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	err = service.RegisterUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Usuário criado com sucesso",
		"user":    user,
	})
}
