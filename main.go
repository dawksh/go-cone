package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pinecone-io/go-pinecone/pinecone"
)

func prettifyStruct(obj interface{}) string {
	bytes, _ := json.MarshalIndent(obj, "", "  ")
	return string(bytes)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ctx := context.Background()

	pc, err := pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: os.Getenv("PINECONE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create Client: %v", err)
	}

	indexName := "index"

	idx, err := pc.CreateServerlessIndex(ctx, &pinecone.CreateServerlessIndexRequest{
		Name:      indexName,
		Dimension: 8,
		Metric:    pinecone.Cosine,
		Cloud:     pinecone.Aws,
		Region:    "us-east-1",
	})

	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	log.Printf("Index created: %v", prettifyStruct(idx))

}
