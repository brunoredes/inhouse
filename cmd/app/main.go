package main

import (
	"ayaxos-inhouse/config/database"
	"ayaxos-inhouse/internal/routes"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	Init_Database()

	r := gin.Default()
	routes.SetupRoutes(r) // Register all routes

	r.Run(":8080") // Start the server on port 8080

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	fmt.Println("🚀 Server is running... Press Ctrl+C to exit.")

	// Wait for termination signal
	<-quit
	fmt.Println("\n🔴 Server shutting down...")
}

func Init_Database() {
	database.InitRedis() // Initialize Redis
	database.ConnectDB() // Initialize PG connection

	// defer database.CloseDB()

}
