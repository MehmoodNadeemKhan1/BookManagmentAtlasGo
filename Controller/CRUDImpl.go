package controller

import (
	"fmt"

	"github.com/Mehmoodkhan1/BookStore/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type NewCustomerImpl struct {
	DB *gorm.DB
}

func NewCustomerControllerImpl(db *gorm.DB) Crud {
	return &NewCustomerImpl{DB: db}
}

// CreateCustomer implements Crud.
func (n *NewCustomerImpl) CreateCustomer(c *fiber.Ctx) error {
	var customer models.Customer
	if err := c.BodyParser(&customer); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := n.DB.Create(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Println("Customer created successfully:", customer)
	return c.Status(fiber.StatusCreated).JSON(customer)
}

// DeleteCustomer implements Crud.
func (n *NewCustomerImpl) DeleteCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer models.Customer
	if err := n.DB.First(&customer, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}

	if err := n.DB.Delete(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Customer deleted"})
}

// GetCustomer implements Crud.
func (n *NewCustomerImpl) GetCustomer(c *fiber.Ctx) error {
	id := c.Params("id")
	var customer models.Customer
	if err := n.DB.Preload("Address").Preload("PhoneNumbers").First(&customer, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}
	return c.JSON(customer)
}

// GetCustomers implements Crud.
func (n *NewCustomerImpl) GetCustomers(c *fiber.Ctx) error {
	var customers []models.Customer
	if err := n.DB.Preload("Address").Preload("PhoneNumbers").Find(&customers).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(customers)
}

// UpdateCustomer implements Crud.
func (n *NewCustomerImpl) UpdateCustomer(c *fiber.Ctx) error {
	id := c.Params("id")

	var customer models.Customer
	if err := n.DB.Preload("Address").Preload("PhoneNumbers").First(&customer, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Customer not found"})
	}

	var updated models.Customer
	if err := c.BodyParser(&updated); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	// Update basic customer fields
	customer.CustomerName = updated.CustomerName
	customer.CustomerEmail = updated.CustomerEmail

	if err := n.DB.Save(&customer).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	for _, updatedAddr := range updated.Address {
		err := n.DB.Model(&models.CustomerAddress{}).
			Where("address_id = ? AND customer_id = ?", updatedAddr.AddressID, customer.ID).
			Update("address", updatedAddr.Address).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to update address ID %d: %v", updatedAddr.AddressID, err),
			})
		}
	}

	for _, updatedPhone := range updated.PhoneNumbers {
		err := n.DB.Model(&models.CustomerPhoneNumber{}).
			Where("phone_number_id = ? AND customer_id = ?", updatedPhone.PhoneNumberID, customer.ID).
			Update("phone_number", updatedPhone.PhoneNumber).Error
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": fmt.Sprintf("Failed to update phone number ID %d: %v", updatedPhone.PhoneNumberID, err),
			})
		}
	}

	if err := n.DB.Preload("Address").Preload("PhoneNumbers").First(&customer, id).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching updated customer"})
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}
