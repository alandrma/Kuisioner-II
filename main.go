package main

import (
	"Kuisioner-MySql/db"
	"Kuisioner-MySql/models"

	"Kuisioner-MySql/routes"
)

func main() {

	db := db.SetupDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Kuisioner{})

	r := routes.SetupRoutes(db)
	r.Run(":8181")
}
