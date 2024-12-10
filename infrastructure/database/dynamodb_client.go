package database

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// DynamoDBClient contiene el cliente para interactuar con DynamoDB
type DynamoDBClient struct {
	Client *dynamodb.Client
	Table  string
}

// NewDynamoDBClient inicializa y devuelve un cliente de DynamoDB
func NewDynamoDBClient(tableName string) *DynamoDBClient {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		log.Fatalf("AWS_REGION environment variable is not set")
	}

	// Configurar AWS SDK con la región obtenida
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("Failed to load AWS configuration: %v", err)
	}

	return &DynamoDBClient{
		Client: dynamodb.NewFromConfig(cfg),
		Table:  tableName,
	}
}
