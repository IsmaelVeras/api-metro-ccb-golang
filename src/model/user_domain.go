package model

import (
	"crypto/md5"
	"encoding/hex"

	rest_err "github.com/IsmaelVeras/api-golang-crud/src/configuration"
)

func NewUserDomain(
	email, password, name string,
	age int8,
) *UserDomain {
	return &UserDomain{
		email, password, name, age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) (UserDomain, *rest_err.RestErr)
	DeleteUser(string) *rest_err.RestErr
}
