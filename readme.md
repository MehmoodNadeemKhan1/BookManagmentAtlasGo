**ATLAS MIGRATION TOOL**
curl -sSf https://atlasgo.sh | sh
**go library that read go struct and automate generate the sql schema**
go get -u ariga.io/atlas-provider-gorm
1. download all libraries given above for migration 
2. make atlas.hcl file 

data "external_schema" "gorm" {
  program = [
    "go", "run", "./loader"
  ]
}

env "gorm" {
 src = data.external_schema.gorm.url
  dev = "postgres://postgres:9009@localhost:5432/BookStoreDB?sslmode=disable"
  url = "postgres://postgres:9009@localhost:5432/BookStoreDB?sslmode=disable"

  migration {
    dir = "file://migrations"
  }
}

put your database configuration 

3. make loader name folder make main.go file in and put the code 


**Atlas migration commands**
atlas migrate diff add_products_table --env gorm
atlas migrate apply --env gorm
