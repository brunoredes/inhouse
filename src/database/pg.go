package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib" // pgx driver
	"github.com/joho/godotenv"
)

// DB is a global variable to access the database connection.
var DB *pgx.Conn

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

	fmt.Println(dsn)

	// Create a background context
	ctx := context.Background()

	// Connect to the database
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Assign connection to global DB variable
	DB = conn

	fmt.Println("✅ Connected to PostgreSQL successfully")
}

// CloseDB closes the database connection
func CloseDB() {
	if DB != nil {
		DB.Close(context.Background())
		fmt.Println("🛑 PostgreSQL connection closed")
	}
}
