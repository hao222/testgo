package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth(
		"vwxdzufqorjdjaac",
		"1154714871",
		"",
		"smtp.qq.com",
	)
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(
		"smtp.qq.com:25",
		auth,
		"1154714871@qq.com",
		[]string{"1154714871@qq.com"},
		[]byte("This is the email body."),
	)
	if err != nil {
		log.Fatal(err)
	}
}
