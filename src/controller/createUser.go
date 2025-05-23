package controller

import (
	"github.com/IsmaelVeras/api-golang-crud/src/configuration/logger"
	"github.com/IsmaelVeras/api-golang-crud/src/configuration/validation"
	"github.com/IsmaelVeras/api-golang-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Iniciando o controller CreateUser",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Erro ao fazer o bind do JSON", err)
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}
	logger.Info("User Criado com sucesso",
		zap.String("journey", "createUser"),
	)
}
