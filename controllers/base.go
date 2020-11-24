package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type Controller interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Delete(c *gin.Context)
	Store(c *gin.Context)
}
