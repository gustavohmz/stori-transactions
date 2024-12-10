package interfaces

import "stori-transactions/domain/models"

// TransactionRepository define los métodos para interactuar con las transacciones en la base de datos
type TransactionRepository interface {
	SaveTransaction(transaction models.Transaction) error
	GetAllTransactions(accountID string) ([]models.Transaction, error)
}
