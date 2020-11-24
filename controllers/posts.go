package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"api/models"
)

type PostEntity interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Delete(c *gin.Context)
	Store(c *gin.Context)
}

type PostsCtrl struct{}

var postModel = models.Post{}

func (ctrl *PostsCtrl) Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "world",
	})
}

func (ctrl *PostsCtrl) Index(c *gin.Context) {
	posts, _ := postModel.All()
	c.JSON(http.StatusOK, posts)
}

func (ctrl *PostsCtrl) Show(c *gin.Context) {
	id := c.Params.ByName("post")

	post := postModel.Show(id)
	c.JSON(http.StatusOK, post)
}

func (ctrl *PostsCtrl) Delete(c *gin.Context) {
	id := c.Params.ByName("post")

	postModel.Delete(id)
	c.JSON(http.StatusOK, "successfully deleted")
}

func (ctrl *PostsCtrl) Store(c *gin.Context) {
	var post models.Post
	c.BindJSON(&post)

	validate = validator.New()
	if err := validate.Struct(post); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	post.UserId = 1

	postModel.Create(&post)
	c.JSON(http.StatusOK, post)
}
