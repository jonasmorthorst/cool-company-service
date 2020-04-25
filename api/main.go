package main

import (
	"api/app"
)

func main() {
	app := &app.App{}
	app.Initialize()
	app.Run(":8080")
}
