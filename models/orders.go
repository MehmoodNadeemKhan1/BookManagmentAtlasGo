package models

import "gorm.io/gorm"

type Orders struct {
	gorm.Model                // adds ID, CreatedAt, UpdatedAt, DeletedAt
	CustomerID    int         `json:"customer_id" gorm:"not null"`
	PaymentMethod string      `json:"payment_method" gorm:"not null"`
	OrderDate     string      `json:"order_date" gorm:"not null"`
	OrderItem     []OrderItem `json:"order_item" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID  uint    `json:"order_id" gorm:"not null;index"` // Add index for foreign key
	Quantity int     `json:"quantity" gorm:"not null"`
	Price    float64 `json:"price" gorm:"not null"`
	Books    []Books `json:"books" gorm:"foreignKey:BookISBN;"`
}
