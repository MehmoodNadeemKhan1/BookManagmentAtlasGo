package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	CustomerName  string                `json:"customer_name" gorm:"not null"`
	CustomerEmail string                `json:"customer_email" gorm:"not null;unique"`
	Address       []CustomerAddress     `json:"address" gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE"`
	PhoneNumbers  []CustomerPhoneNumber `json:"phone_numbers" gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE"`
	Orders        []Orders              `json:"orders" gorm:"foreignKey:CustomerID;constraint:OnDelete:CASCADE"`
}

type CustomerAddress struct {
	AddressID  int    `json:"address_id" gorm:"primaryKey"`
	CustomerID uint   `json:"customer_id" gorm:"not null"`
	Address    string `json:"street_address" gorm:"column:address;not null"` // FIXED: maps "street_address" to actual "address" column
}

type CustomerPhoneNumber struct {
	PhoneNumberID int    `json:"phone_number_id" gorm:"primaryKey"`
	CustomerID    uint   `json:"customer_id" gorm:"not null"`
	PhoneNumber   string `json:"phone_number" gorm:"not null;unique"`
}
