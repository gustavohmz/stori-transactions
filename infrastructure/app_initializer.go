package infrastructure

import (
	"log"
	"os"
	"stori-transactions/domain/interfaces"
	"stori-transactions/domain/services"
	"stori-transactions/infrastructure/aws"
	"stori-transactions/infrastructure/database"
	"stori-transactions/infrastructure/email"
)

type AppInitializer struct {
	TransactionRepo    interfaces.TransactionRepository
	TransactionService interfaces.TransactionProcessor
	SummaryService     interfaces.SummaryCalculator
	EmailSender        interfaces.EmailSender
	S3Client           *aws.S3Client
}

func NewAppInitializer() *AppInitializer {
	// Inicializar DynamoDB
	tableName := os.Getenv("DYNAMODB_TABLE")
	if tableName == "" {
		log.Fatalf("DYNAMODB_TABLE environment variable is not set")
	}
	dynamoClient := database.NewDynamoDBClient(tableName)
	transactionRepo := database.NewDynamoDBTransactionRepo(dynamoClient)

	// Inicializar S3
	bucketName := os.Getenv("S3_BUCKET")
	if bucketName == "" {
		log.Fatalf("S3_BUCKET environment variable is not set")
	}
	s3Client := aws.NewS3Client(bucketName)

	// Inicializar servicios
	transactionService := services.NewTransactionService(transactionRepo)
	summaryService := services.NewSummaryService()
	emailService := email.NewEmailService()

	return &AppInitializer{
		TransactionRepo:    transactionRepo,
		TransactionService: transactionService,
		SummaryService:     summaryService,
		EmailSender:        emailService,
		S3Client:           s3Client,
	}
}
