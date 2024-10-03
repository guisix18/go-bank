package main

import (
	"go-bank/config"
	controllers "go-bank/controllers/user"
	"go-bank/services/user"

	"github.com/gin-gonic/gin"
)

func main() {
    // Conectar ao banco de dados
    config.ConnectDatabase()

    // Inicializar o router Gin
    router := gin.Default()

    // Criar uma instância do UserController
    userController := controllers.UserController{
        Service: user.UserService{},
    }

    // Definir a rota para obter todos os usuários
    router.GET("/users", userController.GetAllUsers)

    // Iniciar o servidor
    router.Run(":8080")
}
