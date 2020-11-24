package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"

	"api/libs"
	"api/models"
	"api/routes"
)

func init() {
	libs.InitLogrus()

	if err := godotenv.Load(); err != nil {
		log.Error(err)
	}

	if err := models.InitDB(); err != nil {
		log.Error(err)
		return
	}

	libs.InitRedis()
}

func main() {
	routes.Init()

	routes.Router.Run(":5000")
}
