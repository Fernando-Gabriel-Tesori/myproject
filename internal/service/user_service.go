package service

import (
	"errors"
	"strings"

	"myproject/internal/repository"
	"myproject/models"
)

func RegisterUser(user *models.User) error {

	if strings.TrimSpace(user.Name) == "" || strings.TrimSpace(user.Email) == "" || strings.TrimSpace(user.Password) == "" {
		return errors.New("todos os campos são obrigatórios")
	}

	user.Email = strings.ToLower(strings.TrimSpace(user.Email))

	existing, _ := repository.GetUserByEmail(user.Email)
	if existing != nil {
		return errors.New("email já cadastrado")
	}

	return repository.CreateUser(user)
}
