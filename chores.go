package main

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendEmail(from, password string, message byte) {
	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	err := smtp.SendMail(
		"smtp.gmail.com",
		auth,
		from,
		[]string{"shapiroj18@gmail.com"},
		[]byte{message},
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
	goDotEnvVariable("STRONGEST_AVENGER")
	// SendEmail()
}
