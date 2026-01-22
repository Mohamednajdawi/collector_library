package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"collector-library/internal/adapters/repository"
	"collector-library/internal/core/domain"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	// Connect to Database
	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer dbPool.Close()

	repo := repository.NewPostgresRepository(dbPool)

	// Read JSON data
	dataPath := filepath.Join("..", "..", "data", "amiibo_raw.json")
	file, err := os.ReadFile(dataPath)
	if err != nil {
		// Try current directory if relative fails
		file, err = os.ReadFile("data/amiibo_raw.json")
		if err != nil {
			log.Fatalf("Failed to read amiibo data: %v", err)
		}
	}

	var amiibos []domain.Amiibo
	if err := json.Unmarshal(file, &amiibos); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	fmt.Printf("found %d amiibos to ingest...\n", len(amiibos))

	// Ingest
	ctx := context.Background()
	if err := repo.CreateBatch(ctx, amiibos); err != nil {
		log.Fatalf("Failed to ingest data: %v", err)
	}

	fmt.Println("Successfully ingested all Amiibos!")
}
