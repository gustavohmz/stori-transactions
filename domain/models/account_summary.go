package models

// AccountSummary contiene un resumen del estado de una cuenta
type AccountSummary struct {
	Balance          float64                 `json:"balance"`
	MonthlySummaries map[string]MonthlyStats `json:"monthly_summaries"`
}
