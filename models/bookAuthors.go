package models

import "gorm.io/gorm"

type Books struct {
	BookISBN  int       `gorm:"primaryKey;autoIncrement"`
	BookTitle string    `gorm:"not null"`
	BookPrice float64   `gorm:"not null"`
	Authors   []Authors `gorm:"many2many:book_authors;"`
}

type Authors struct {
	AuthorId    int     `gorm:"primaryKey;autoIncrement"`
	AuthorName  string  `gorm:"not null"`
	AuthorEmail string  `gorm:"not null;unique"`
	AuthorBio   string  `gorm:"not null"`
	Books       []Books `gorm:"many2many:book_authors;"`
}

type BookAuthors struct {
	gorm.Model
	BookISBN int `gorm:"primaryKey"`
	AuthorId int `gorm:"primaryKey"`
}
