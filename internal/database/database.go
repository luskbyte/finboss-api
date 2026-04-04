package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

func Connect() *sql.DB {
	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		log.Fatal("DB_PASSWORD environment variable is required")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/San_Paulo",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "finboss"),
		password,
		getEnv("DB_NAME", "finboss"),
		getEnv("DB_PORT", "5433"),
		getEnv("DB_SSLMODE", "require"),
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("database connected")
	return db
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
