package user

import (
	"go-bank/config"
	"go-bank/models"
)

// UserService é a struct que contém métodos relacionados a usuários.
type UserService struct{}

// GetAllUsers retorna todos os usuários do banco de dados.
func (s *UserService) GetAllUsers() ([]models.User, error) {
    var users []models.User
    if err := config.DB.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}
