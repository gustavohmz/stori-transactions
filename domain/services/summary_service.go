package services

import "stori-transactions/domain/models"

// SummaryService implementa SummaryCalculator
type SummaryService struct{}

// NewSummaryService crea una nueva instancia de SummaryService
func NewSummaryService() *SummaryService {
	return &SummaryService{}
}

// CalculateSummary calcula el resumen de transacciones
func (s *SummaryService) CalculateSummary(transactions []models.Transaction) models.AccountSummary {
	balance := 0.0
	monthlyStats := map[string]models.MonthlyStats{}

	for _, tx := range transactions {
		balance += tx.Amount
		month := tx.Date.Format("January 2006")

		if _, exists := monthlyStats[month]; !exists {
			monthlyStats[month] = models.MonthlyStats{}
		}

		stats := monthlyStats[month]
		stats.TotalTransactions++

		if tx.Type == "credit" {
			stats.AverageCredit += tx.Amount
		} else if tx.Type == "debit" {
			stats.AverageDebit += tx.Amount
		}

		monthlyStats[month] = stats
	}

	for month, stats := range monthlyStats {
		stats.AverageCredit /= float64(stats.TotalTransactions)
		stats.AverageDebit /= float64(stats.TotalTransactions)
		monthlyStats[month] = stats
	}

	return models.AccountSummary{
		Balance:          balance,
		MonthlySummaries: monthlyStats,
	}
}
