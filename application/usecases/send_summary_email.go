package usecases

import (
	"stori-transactions/domain/interfaces"
	"stori-transactions/domain/models"
)

// SendSummaryEmail genera y envía un correo basado en el resumen
func SendSummaryEmail(
	emailSender interfaces.EmailSender, // Servicio que implementa la interfaz
	summary models.AccountSummary, // Resumen de las transacciones
	recipient string, // Dirección de correo del destinatario
) error {
	// Delegar al servicio que implemente la interfaz
	return emailSender.SendSummaryEmail(summary, recipient)
}
