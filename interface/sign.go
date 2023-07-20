package intf

import (
	"main/utils"
	"strings"
	"github.com/gin-gonic/gin"
)

type SignInputD struct {
	E_mail string `json:"e-mail" form:"e-mail" yaml:"e-mail"`
	Passwd string `json:"passwd" form:"passwd" yaml:"passwd"`
	SignInput
}

type SignInput struct {
	Email    string `json:"email" form:"email" yaml:"email"`
	Password string `json:"password" form:"password" yaml:"password"`
}

// Get SignInput from context
func ContextSignInput(c *gin.Context) (SignInput, error) {
	dataD := SignInputD{}
	err := utils.BindWithHeader(c, &dataD)
	data := SignInputReduce(dataD)
	return data, err
}

// Reduce SignInputD to SignInput
func SignInputReduce(s SignInputD) SignInput {
	signin := SignInput{
		Email:    strings.ToLower(s.Email),
		Password: strings.ToLower(s.Password),
	}
	if signin.Email == "" {
		signin.Email = strings.ToLower(s.E_mail)
	}
	if signin.Password == "" {
		signin.Password = strings.ToLower(s.Passwd)
	}
	return signin
}
