package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	dbPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Unable to create database pool: %v\n", err)
	}
	defer dbPool.Close()

	// Direct Create Table Query
	// We skip Policies in this fallback script to ensure the table exists for seeding.
	query := `
	CREATE TABLE IF NOT EXISTS public.amiibos (
		id uuid not null default gen_random_uuid (),
		name text not null,
		image_url text null,
		series text not null default 'Super Smash Bros',
		release_date date null,
		created_at timestamp with time zone not null default now(),
		constraint amiibos_pkey primary key (id)
	);
	`

	log.Println("Creating table if not exists...")
	if _, err := dbPool.Exec(context.Background(), query); err != nil {
		log.Fatalf("Failed to execute create table: %v", err)
	}

	log.Println("Table 'public.amiibos' is ready!")
}
