package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"

	"api/apis"
	"api/libs"
	"api/models"
)

type UsersCtrl struct{}

var userModel = models.User{}

func (ctrl *UsersCtrl) Index(c *gin.Context) {
	uParams := c.MustGet("userParams").(*libs.TokenData)
	// var waitGroup sync.WaitGroup
	// waitGroup.Add(2)
	// defer waitGroup.Wait()

	apis.Get()

	users := userModel.All(uParams.TeamId)
	c.JSON(http.StatusOK, users)
}

func (ctrl *UsersCtrl) Show(c *gin.Context) {
	id := c.Params.ByName("user")

	user := userModel.Show(id)
	c.JSON(http.StatusOK, user)
}

func (ctrl *UsersCtrl) Delete(c *gin.Context) {
	id := c.Params.ByName("user")

	userModel.Delete(id)
	c.JSON(http.StatusOK, "successfully deleted")
}

func (ctrl *UsersCtrl) Store(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	validate = validator.New()
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userModel.Create(&user)
	c.JSON(http.StatusOK, user)
}

func (ctrl *UsersCtrl) Csv(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Error(err)
		return
	}
	filename := header.Filename
	out, err := os.Create("tmp/" + filename)
	if err != nil {
		log.Error(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Error(err)
	}

	c.JSON(http.StatusOK, nil)
}
