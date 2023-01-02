package main

import (
	"go_postgres/initializers"
	"go_postgres/models"
)

func init() {
	initializers.LoadEnvVvariables()
	initializers.ConnectToDB()

}

func main() {
	initializers.DB.AutoMigrate(&models.Student{})
}
