package interfaces

import "stori-transactions/domain/models"

// EmailSender define la interfaz para enviar correos electrónicos
type EmailSender interface {
	SendSummaryEmail(summary models.AccountSummary, recipient string) error
}
