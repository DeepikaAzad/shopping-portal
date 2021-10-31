package main

import (
	"fmt"

	"github.com/DeepikaAzad/go-to-do-app/go-server/app"
)

func main() {
	fmt.Println("Starting server on the port 8080...")
	app.StartApplication()
}
