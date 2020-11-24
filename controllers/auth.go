package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"api/libs"
	"api/models"
)

type AuthCtrl struct{}

func (ctrl *AuthCtrl) Login(c *gin.Context) {
	var login models.Login
	c.BindJSON(&login)

	validate = validator.New()
	if err := validate.Struct(login); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := libs.SignIn(12, 12)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	c.JSON(http.StatusOK, token)
}
