package aws

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3Client maneja la interacción con S3
type S3Client struct {
	Client *s3.Client
	Bucket string
}

// NewS3Client crea un nuevo cliente S3
func NewS3Client(bucket string) *S3Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
	if err != nil {
		log.Fatalf("Failed to load AWS configuration: %v", err)
	}

	return &S3Client{
		Client: s3.NewFromConfig(cfg),
		Bucket: bucket,
	}
}

// UploadFile sube un archivo al bucket de S3
func (s *S3Client) UploadFile(fileName string, fileContent []byte) error {
	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(fileContent),
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}
	log.Printf("File uploaded successfully to S3: %s/%s", s.Bucket, fileName)
	return nil
}

// DownloadFile descarga un archivo desde el bucket de S3
func (s *S3Client) DownloadFile(fileName string) ([]byte, error) {
	output, err := s.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file from S3: %w", err)
	}
	defer output.Body.Close()

	// Leer el contenido del archivo
	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, output.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read file content: %w", err)
	}
	return buffer.Bytes(), nil
}
