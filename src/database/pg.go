package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // pgx driver
	"github.com/joho/godotenv"
)

// DB is a global variable to access the database connection.
var DB *sql.DB

// ConnectDB initializes a PostgreSQL database connection.
func ConnectDB() {
	_ = godotenv.Load()

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_NAME"),
		os.Getenv("PG_SSL"),
	)

	// Open a database connection
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Ping database to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Database ping failed: %v", err)
	}

	fmt.Println("✅ Connected to PostgreSQL successfully!")
	DB = db
}

// CloseDB closes the database connection.
func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("🔌 Database connection closed.")
	}
}
