package model

import (
	"fmt"

	rest_err "github.com/IsmaelVeras/api-golang-crud/src/configuration"
	"github.com/IsmaelVeras/api-golang-crud/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init CreateUser model", zap.String("journey", "CreateUser"))

	ud.EncryptPassword()
	fmt.Println(ud)
	return nil
}
