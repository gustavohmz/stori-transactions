package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"stori-transactions/domain/models"
)

// EmailService implementa la interfaz EmailSender
type EmailService struct {
	from     string
	username string
	password string
	smtpHost string
	smtpPort string
	logoURL  string
}

// NewEmailService crea una nueva instancia del servicio de email
func NewEmailService() *EmailService {
	return &EmailService{
		from:     os.Getenv("SMTP_FROM"),
		username: os.Getenv("SMTP_USERNAME"),
		password: os.Getenv("SMTP_PASSWORD"),
		smtpHost: os.Getenv("SMTP_HOST"),
		smtpPort: os.Getenv("SMTP_PORT"),
		logoURL:  os.Getenv("EMAIL_LOGO_URL"),
	}
}

// SendSummaryEmail implementa la interfaz para enviar un correo basado en el resumen
func (e *EmailService) SendSummaryEmail(summary models.AccountSummary, recipient string) error {
	// Cargar la plantilla HTML del correo
	tmpl, err := template.ParseFiles("templates/email_template.html")
	if err != nil {
		log.Printf("ERROR: Failed to load email template: %v", err)
		return err
	}

	// Datos para inyectar en el template
	data := struct {
		LogoURL          string
		TotalBalance     float64
		MonthlySummaries map[string]models.MonthlyStats
	}{
		LogoURL:          e.logoURL,
		TotalBalance:     summary.Balance,
		MonthlySummaries: summary.MonthlySummaries,
	}

	// Procesar el template con los datos
	var htmlBody bytes.Buffer
	err = tmpl.Execute(&htmlBody, data)
	if err != nil {
		log.Printf("ERROR: Failed to execute email template: %v", err)
		return err
	}

	// Configuración del correo con asunto
	subject := "Transaction Summary"
	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s",
		e.from, recipient, subject, htmlBody.String(),
	)

	// Autenticación SMTP
	auth := smtp.PlainAuth("", e.username, e.password, e.smtpHost)

	// Enviar el correo
	err = smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.from, []string{recipient}, []byte(msg))
	if err != nil {
		log.Printf("ERROR: Failed to send email using SMTP: %v", err)
		return err
	}

	log.Println("Email sent successfully")
	return nil
}
