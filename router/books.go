package router

import (
	"github.com/aks1809/librarymanagement_backend/controller"
	"github.com/gofiber/fiber/v2"
)

func BookRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	api.Get("/books", controller.GetBooks)
	api.Get("/book/:id", controller.GetBook)
	api.Post("/book", controller.CreateBook)
	api.Put("/book/:id", controller.UpdateBook)
	api.Delete("/book/:id", controller.DeleteBook)
}
