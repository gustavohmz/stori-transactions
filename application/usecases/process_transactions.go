package usecases

import (
	"stori-transactions/domain/interfaces"
	"stori-transactions/domain/models"
)

// ProcessTransactions valida, guarda y calcula resúmenes
func ProcessTransactions(
	processor interfaces.TransactionProcessor,
	repo interfaces.TransactionRepository,
	calculator interfaces.SummaryCalculator,
	transactions []models.Transaction,
) (models.AccountSummary, error) {
	for _, tx := range transactions {
		if err := processor.ValidateTransaction(tx); err != nil {
			return models.AccountSummary{}, err
		}
		if err := repo.SaveTransaction(tx); err != nil {
			return models.AccountSummary{}, err
		}
	}
	return calculator.CalculateSummary(transactions), nil
}
