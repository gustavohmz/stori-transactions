package interfaces

import "stori-transactions/domain/models"

// TransactionProcessor define el contrato para procesar transacciones
type TransactionProcessor interface {
	ValidateTransaction(transaction models.Transaction) error
	SaveTransaction(transaction models.Transaction) error
}
