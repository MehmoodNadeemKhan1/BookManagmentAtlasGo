package main

import (
	"fmt"
	"log"

	"ariga.io/atlas-provider-gorm/gormschema"
	"github.com/Mehmoodkhan1/BookStore/models"
)

func main() {
	ddl, err := gormschema.New("postgres").Load(
		&models.Customer{},
		&models.CustomerAddress{},
		&models.CustomerPhoneNumber{},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Print the generated DDL (Atlas will parse this)
	fmt.Print(ddl)
}
