package database

import (
	"context"
	"fmt"
	"log"
	"stori-transactions/domain/models"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoDBTransactionRepo implementa TransactionRepository para DynamoDB
type DynamoDBTransactionRepo struct {
	client *DynamoDBClient
}

// NewDynamoDBTransactionRepo crea un nuevo repositorio de transacciones
func NewDynamoDBTransactionRepo(client *DynamoDBClient) *DynamoDBTransactionRepo {
	return &DynamoDBTransactionRepo{
		client: client,
	}
}

// SaveTransaction guarda una transacción en DynamoDB
func (repo *DynamoDBTransactionRepo) SaveTransaction(transaction models.Transaction) error {
	_, err := repo.client.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(repo.client.Table),
		Item: map[string]types.AttributeValue{
			"ID":        &types.AttributeValueMemberS{Value: transaction.ID},
			"Amount":    &types.AttributeValueMemberN{Value: fmt.Sprintf("%.2f", transaction.Amount)},
			"Type":      &types.AttributeValueMemberS{Value: transaction.Type},
			"Date":      &types.AttributeValueMemberS{Value: transaction.Date.Format("2006-01-02")},
			"AccountID": &types.AttributeValueMemberS{Value: transaction.AccountID},
		},
	})
	if err != nil {
		log.Printf("Failed to save transaction: %v", err)
		return err
	}
	return nil
}

func (repo *DynamoDBTransactionRepo) GetAllTransactions(id string) ([]models.Transaction, error) {
	result, err := repo.client.Client.Query(context.TODO(), &dynamodb.QueryInput{
		TableName:              aws.String(repo.client.Table),
		KeyConditionExpression: aws.String("ID = :id"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		log.Printf("Failed to get transactions: %v", err)
		return nil, err
	}

	var transactions []models.Transaction
	for _, item := range result.Items {
		transaction := models.Transaction{
			ID:        item["ID"].(*types.AttributeValueMemberS).Value,
			Amount:    parseFloat(item["Amount"].(*types.AttributeValueMemberN).Value),
			Type:      item["Type"].(*types.AttributeValueMemberS).Value,
			AccountID: item["AccountID"].(*types.AttributeValueMemberS).Value,
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func parseFloat(value string) float64 {
	parsed, _ := strconv.ParseFloat(value, 64)
	return parsed
}
