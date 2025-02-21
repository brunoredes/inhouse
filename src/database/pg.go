package database

import (
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

	conn, err := pgx.Connect(Ctx, dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	defer conn.Close(Ctx)

	DB = conn
}
