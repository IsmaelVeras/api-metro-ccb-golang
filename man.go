package main

import (
	"log"

	"github.com/IsmaelVeras/api-golang-crud/src/configuration/logger"
	"github.com/IsmaelVeras/api-golang-crud/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Iniciando o servidor...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erro ao inicializar o servidor : ", err)
	}
}
