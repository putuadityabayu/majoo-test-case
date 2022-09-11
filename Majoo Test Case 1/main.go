package main

import (
	"majoo.id/satu/config"
	"majoo.id/satu/routes"
)

func main() {
	// Setup database
	config.SetupDatabase()

	// Setup gofiber
	r := routes.SetupRouters()

	// Run web server
	r.Listen(":3000")
}
