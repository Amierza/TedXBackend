package utils

import (
	"log"
	"path/filepath"

	"github.com/Amierza/TedXBackend/config"
	"gopkg.in/gomail.v2"
)

func SendEmail(toEmail string, subject string, body string) error {
	emailConfig, err := config.NewEmailConfig()
	if err != nil {
		log.Printf("failed to load email config: %v", err)
		return err
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", emailConfig.AuthEmail)
	mailer.SetHeader("To", toEmail)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/html", body)

	// 🔹 Buat path absolut agar aman dijalankan dari mana pun
	assetPath, err := filepath.Abs("assets_static/header-e-ticket-mail.png")
	if err != nil {
		log.Printf("failed to get absolute path: %v", err)
		return err
	}

	// Embed gambar tanpa menangkap nilai return (karena tidak ada)
	mailer.Embed(assetPath)

	dialer := gomail.NewDialer(
		emailConfig.Host,
		emailConfig.Port,
		emailConfig.AuthEmail,
		emailConfig.AuthPassword,
	)

	if err := dialer.DialAndSend(mailer); err != nil {
		log.Printf("failed to send email to %v: %v", toEmail, err)
		return err
	}

	return nil
}
