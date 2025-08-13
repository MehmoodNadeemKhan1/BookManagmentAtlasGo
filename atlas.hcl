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
