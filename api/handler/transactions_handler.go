package handler

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"stori-transactions/application/usecases"
	"stori-transactions/domain/models"
	"stori-transactions/infrastructure"
)

func HandleRequest(ctx context.Context, event map[string]interface{}) error {
	app := infrastructure.NewAppInitializer()

	// Leer el nombre del archivo desde las variables de entorno
	fileName := os.Getenv("S3_INPUT_FILE")
	if fileName == "" {
		log.Println("S3_INPUT_FILE is not set in the environment variables")
		return fmt.Errorf("missing S3_INPUT_FILE environment variable")
	}

	// Paso 1: Descargar el archivo desde S3
	fileContent, err := app.S3Client.DownloadFile(fileName)
	if err != nil {
		log.Printf("Failed to download file from S3: %v\n", err)
		return err
	}

	// Paso 2: Procesar el contenido del archivo
	transactions, err := parseCSVContent(fileContent)
	if err != nil {
		log.Printf("Failed to parse CSV content: %v\n", err)
		return err
	}

	// Paso 3: Validar, guardar y calcular el resumen
	summary, err := usecases.ProcessTransactions(
		app.TransactionService,
		app.TransactionRepo,
		app.SummaryService,
		transactions,
	)
	if err != nil {
		log.Printf("Failed to process transactions: %v\n", err)
		return err
	}

	// Paso 4: Subir el archivo procesado a S3
	csvData := new(bytes.Buffer)
	writer := csv.NewWriter(csvData)

	writer.Write([]string{"ID", "Amount", "Type", "Date", "AccountID"})
	for _, tx := range transactions {
		writer.Write([]string{
			tx.ID,
			fmt.Sprintf("%.2f", tx.Amount),
			tx.Type,
			tx.Date.Format("2006-01-02"),
			tx.AccountID,
		})
	}
	writer.Flush()

	processedFileName := "processed-transactions-" + time.Now().Format("20060102150405") + ".csv"
	err = app.S3Client.UploadFile(processedFileName, csvData.Bytes())
	if err != nil {
		log.Printf("Failed to upload processed file to S3: %v\n", err)
		return err
	}

	// Paso 5: Enviar correo con el resumen
	recipient := os.Getenv("SMTP_TO")
	err = usecases.SendSummaryEmail(app.EmailSender, summary, recipient)
	if err != nil {
		log.Printf("Failed to send email: %v\n", err)
		return err
	}

	log.Println("Handler executed successfully")
	return nil
}

// parseCSVContent analiza el contenido del archivo CSV en memoria
func parseCSVContent(fileContent []byte) ([]models.Transaction, error) {
	reader := csv.NewReader(bytes.NewReader(fileContent))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions []models.Transaction
	for _, record := range records[1:] { // Saltar la primera fila (encabezados)
		amount, _ := strconv.ParseFloat(record[1], 64)
		date, _ := time.Parse("2006-01-02", record[3])

		transactions = append(transactions, models.Transaction{
			ID:        record[0],
			Amount:    amount,
			Type:      record[2],
			Date:      date,
			AccountID: record[4],
		})
	}

	return transactions, nil
}
