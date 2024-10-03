package controllers

import (
	"go-bank/services/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController é a struct para o controlador de usuários.
type UserController struct {
    Service user.UserService
}

// GetAllUsers é o handler para a rota que retorna todos os usuários.
func (ctrl *UserController) GetAllUsers(c *gin.Context) {
    users, err := ctrl.Service.GetAllUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, users)
}
