package main

import (
	"restapi/initializers"
	"restapi/models"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Student{}, &models.Book{})
}