package libs

import (
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

func InitLogrus() {
	log.AddHook(filename.NewHook())
	log.SetFormatter(&log.JSONFormatter{})
}
