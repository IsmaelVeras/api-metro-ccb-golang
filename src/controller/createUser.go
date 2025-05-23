package controller

import (
	"fmt"

	rest_err "github.com/IsmaelVeras/api-golang-crud/src/configuration"
	"github.com/IsmaelVeras/api-golang-crud/src/controller/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Invalid JSON body: error%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println("Nome:", userRequest.Name)
	fmt.Println("Email:", userRequest.Email)
	fmt.Println("Password:", userRequest.Password)
	fmt.Println("Age:", userRequest.Age)
}
