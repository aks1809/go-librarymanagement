package controller

import (
	"github.com/aks1809/librarymanagement_backend/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Books struct {
	gorm.Model
	ISBN string `json:"isbn" db:"isbn"`
	Name string `json:"name" db:"name"`
	YOP  int64  `json:"yop" db:"yop"`
}

func init() {
	db, err := config.OpenDBConnection()
	if err != nil {
		panic(err)
	}
	DB = db
	DB.AutoMigrate(&Books{})
}

func GetBooks(c *fiber.Ctx) error {
	var books []Books
	DB.Find(&books)
	return c.Status(200).JSON(&books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	var book Books
	DB.Find(&book, id)
	if book.ID == 0 {
		return c.Status(404).SendString("Cannot find specified book")
	}
	return c.Status(200).JSON(&book)
}

func CreateBook(c *fiber.Ctx) error {
	book := new(Books)
	if err := c.BodyParser(book); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&book)
	return c.Status(200).JSON(&book)
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	book := new(Books)
	DB.Find(&book, id)
	if book.ID == 0 {
		return c.Status(404).SendString("Cannot find specified book")
	}
	if err := c.BodyParser(&book); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&book)
	return c.Status(200).SendString("Book updated")
}

func DeleteBook(c *fiber.Ctx) error {
	var book Books
	id := c.Params("id")
	DB.Unscoped().Delete(&book, id)
	return c.Status(200).SendString("Book deleted successfully")
}
