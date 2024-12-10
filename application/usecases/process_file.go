package usecases

import (
	"encoding/csv"
	"os"
	"stori-transactions/domain/models"
	"strconv"
	"time"
)

// ProcessFile procesa un archivo CSV y lo convierte en una lista de transacciones
func ProcessFile(filePath string) ([]models.Transaction, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	for _, record := range records {
		amount, _ := strconv.ParseFloat(record[1], 64)
		date, _ := time.Parse("2006-01-02", record[2])

		transactions = append(transactions, models.Transaction{
			ID:        record[0],
			Amount:    amount,
			Date:      date,
			Type:      record[3],
			AccountID: record[4],
		})
	}

	return transactions, nil
}
