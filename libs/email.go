package libs

import (
	"net/smtp"
	"os"

	log "github.com/sirupsen/logrus"
)

const host = "smtp.gmail.com"
const port = "587"

func SendEmail(to []string, message string) {
	// Sender data.
	from := os.Getenv("GMAIL_FROM")
	password := os.Getenv("GMAIL_PASS")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, host)
	content := []byte(message)
	// Sending email.
	if err := smtp.SendMail(host+":"+port, auth, from, to, content); err != nil {
		log.Error(err)
		return
	}

	log.Info("asdasdas")
}
