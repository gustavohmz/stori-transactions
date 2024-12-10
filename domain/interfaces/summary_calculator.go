package interfaces

import "stori-transactions/domain/models"

// SummaryCalculator define el contrato para calcular resúmenes
type SummaryCalculator interface {
	CalculateSummary(transactions []models.Transaction) models.AccountSummary
}
