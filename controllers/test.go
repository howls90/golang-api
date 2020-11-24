package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestCtrl struct{}

func (ctrl *TestCtrl) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}
