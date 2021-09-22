package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendEmail(from, password string, to []string, message []byte) {
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		from,
		to,
		message,
	)

	if err != nil {
		log.Print(err)
	}
}

func goDotEnvVariable(key string) string {
	err_load := godotenv.Load()

	if err_load != nil {
		log.Fatalf("Error loading .env file")
	}

	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatal("Missing key")
	}

	return val
}

func Adder(a, b int) int {
	return a + b
}

func main() {
	gmail_user := goDotEnvVariable("GMAIL_USER")
	gmail_pass := goDotEnvVariable("GMAIL_PASS")
	jonathan_email, christiana_email := goDotEnvVariable("JONATHAN_EMAIL"), goDotEnvVariable("CHRISTIANA_EMAIL")
	SendEmail(gmail_user, gmail_pass, []string{jonathan_email, christiana_email}, []byte("test email working"))
}
