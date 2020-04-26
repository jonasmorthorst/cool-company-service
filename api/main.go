package main

import (
	"api/app"
	"os"
)

func main() {
	app := &app.App{}

	// Init the db etc
	app.Initialize()

	// Start server on the port given by the env
	port := os.Getenv("PORT")
	app.Run(":" + port)
}
