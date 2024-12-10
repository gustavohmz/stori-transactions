package main

import (
	"context"
	"log"
	"stori-transactions/api/handler"

	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Crear un contexto para la ejecución
	ctx := context.Background()

	// Ejecutar el handler principal
	event := map[string]interface{}{}
	err = handler.HandleRequest(ctx, event)
	if err != nil {
		log.Fatalf("Handler execution failed: %v", err)
	}

	log.Println("Application executed successfully!")
}
