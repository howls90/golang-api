package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"api/models"
)

type CommentsCtrl struct{}

var commentModel = models.Comment{}

func (ctrl *CommentsCtrl) Index(c *gin.Context) {
	teamId := c.MustGet("teamId").(int)

	comments := commentModel.All(teamId)
	c.JSON(http.StatusOK, comments)
}

func (ctrl *CommentsCtrl) Show(c *gin.Context) {
	id := c.Params.ByName("id")

	comment := commentModel.Show(id)
	c.JSON(http.StatusOK, comment)
}

func (ctrl *CommentsCtrl) Delete(c *gin.Context) {
	id := c.Params.ByName("id")

	commentModel.Delete(id)
	c.JSON(http.StatusOK, "successfully deleted")
}

func (ctrl *CommentsCtrl) Store(c *gin.Context) {
	var comment models.Comment
	c.BindJSON(&comment)

	validate = validator.New()
	if err := validate.Struct(comment); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	commentModel.Create(&comment)
	c.JSON(http.StatusOK, comment)
}
