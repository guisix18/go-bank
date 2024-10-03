package config

import (
	"fmt"
	"go-bank/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    // Carregar as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }

    // Pegar as variáveis de ambiente
    user := os.Getenv("POSTGRES_USER")
    password := os.Getenv("POSTGRES_PASSWORD")
    dbname := os.Getenv("POSTGRES_DB")
    port := os.Getenv("POSTGRES_PORT")

   fmt.Print(user)
   fmt.Print(password)
   fmt.Print(dbname)
   fmt.Print(port)

    // Montar o DSN (Data Source Name)
    dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, dbname, port)

    // Conectar ao banco de dados
    database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Falha ao conectar ao banco de dados:", err)
    }

    // Migração automática das tabelas
    database.AutoMigrate(&models.User{}, &models.CodeVerification{})

    DB = database
}
