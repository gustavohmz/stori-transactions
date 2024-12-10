package models

// MonthlyStats representa estadísticas mensuales de una cuenta
type MonthlyStats struct {
	Month             string  `json:"month"`
	TotalTransactions int     `json:"total_transactions"`
	AverageCredit     float64 `json:"average_credit"`
	AverageDebit      float64 `json:"average_debit"`
}
