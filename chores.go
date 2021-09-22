package main

import (
	"fmt"
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

func mainEmail() {
	gmail_user := goDotEnvVariable("GMAIL_USER")
	gmail_pass := goDotEnvVariable("GMAIL_PASS")
	jonathan_email := goDotEnvVariable("JONATHAN_EMAIL")
	christiana_email := goDotEnvVariable("CHRISTIAN_EMAIL")
	sarah_email := goDotEnvVariable("SARAH_EMAIL")
	paul_email := goDotEnvVariable("PAUL_EMAIL")

	google_doc := goDotEnvVariable("CHORES_CHART")

	subject := "Subject: Chores!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(
		`
			<html>
				<body>
					Chores are due today! Please check your name off in the <a href="%s">spreadsheet</a> when you have completed them. Also CC smells like wet feet that have been in a swamp all day.
				</body>
			</html>
		`, google_doc)

	SendEmail(gmail_user, gmail_pass, []string{jonathan_email, christiana_email, sarah_email, paul_email}, []byte(subject+mime+body))
}

func main() {
	mainEmail()
}
