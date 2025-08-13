package controller

import "github.com/gofiber/fiber/v2"

type Crud interface {
	CreateCustomer(c *fiber.Ctx) error
	GetCustomers(c *fiber.Ctx) error
	GetCustomer(c *fiber.Ctx) error
	UpdateCustomer(c *fiber.Ctx) error
	DeleteCustomer(c *fiber.Ctx) error
}
