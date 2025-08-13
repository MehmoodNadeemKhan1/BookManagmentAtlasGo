package config

import (
	"fmt"

	"github.com/Mehmoodkhan1/BookStore/helper"
	"github.com/Mehmoodkhan1/BookStore/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDB(config Config) *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)
	fmt.Println("Database connected successfully")

	// AutoMigrate
	db.AutoMigrate(
		models.Customer{},
		models.CustomerAddress{},
		models.CustomerPhoneNumber{},
		models.Authors{},
		models.Books{},
		models.BookAuthors{},
		models.Orders{},
		models.OrderItem{},
	)
	return db
}
