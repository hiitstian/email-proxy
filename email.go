package main

import (
	"fmt"
	"net/smtp"
)

type EmailConfig struct {
	From     string   `json:"from"`
	Password string   `json:"password"`
	To       []string `json:"to"`
	SmtpHost string   `json:"smtp_host"`
	SmtpPort string   `json:"smtp_port"`
	Subject  string   `json:"subject"`
	Mime     string   `json:"mime"`
	Body     string   `json:"body"`
}

func sendEmail(config EmailConfig) error {

	// Authentication.
	auth := smtp.PlainAuth("", config.From, config.Password, config.SmtpHost)

	fmt.Println(config.SmtpHost + ":" + config.SmtpPort)
	// Sending email.
	message := []byte(
		"Subject: " + config.Subject + "\n" +
			config.Mime + "\n" +
			config.Body)
	err := smtp.SendMail(config.SmtpHost+":"+config.SmtpPort, auth, config.From, config.To, message)
	return err
}
