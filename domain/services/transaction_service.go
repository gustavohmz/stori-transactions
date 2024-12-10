package services

import (
	"errors"
	"stori-transactions/domain/interfaces"
	"stori-transactions/domain/models"
)

// TransactionService implementa TransactionProcessor
type TransactionService struct {
	repo interfaces.TransactionRepository
}

// NewTransactionService crea una nueva instancia de TransactionService
func NewTransactionService(repo interfaces.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

// ValidateTransaction valida una transacción
func (s *TransactionService) ValidateTransaction(transaction models.Transaction) error {
	if transaction.Amount == 0 {
		return errors.New("invalid transaction amount: amount cannot be zero")
	}

	// Validar el tipo de transacción
	if transaction.Type != "credit" && transaction.Type != "debit" {
		return errors.New("invalid transaction type: must be 'credit' or 'debit'")
	}

	return nil
}

// SaveTransaction guarda una transacción
func (s *TransactionService) SaveTransaction(transaction models.Transaction) error {
	// Aquí normalmente interactuarías con un repositorio
	return s.repo.SaveTransaction(transaction)
}
