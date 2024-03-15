package main

import (
	"assignment2/database"
	"assignment2/router"
)

func main() {
	// Initialize the database connection
	database.InitDB()

	r := router.InitRouter()

	// Run the server
	r.Run()
}
