package main

import (
	"log"

	controller "github.com/Mehmoodkhan1/BookStore/Controller"
	"github.com/Mehmoodkhan1/BookStore/config"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	loadConfig, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db := config.ConnectionDB(loadConfig)

	customerController := &controller.NewCustomerImpl{DB: db}

	bookGroup := app.Group("/books")
	bookGroup.Post("/", customerController.CreateCustomer)
	bookGroup.Get("/", customerController.GetCustomers)
	bookGroup.Get("/:id", customerController.GetCustomer)
	bookGroup.Patch("/:id", customerController.UpdateCustomer)
	bookGroup.Delete("/:id", customerController.DeleteCustomer)

	log.Fatal(app.Listen(":3000"))
	log.Println("Server is running on http://localhost:3000")
}
