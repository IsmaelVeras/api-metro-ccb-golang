package controller

import (
	rest_err "github.com/IsmaelVeras/api-golang-crud/src/configuration"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Você chamou a rota de forma errada")
	c.JSON(err.Code, err)
}
