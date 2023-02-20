package main

import (
	"log"
	"os"

	"github.com/aks1809/librarymanagement_backend/router"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.BookRoutes(app)

	log.Fatal(app.Listen(os.Getenv("SERVER_URL")))
}
