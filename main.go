package main

import (
	"MygarmProject/database"
	"MygarmProject/router"
)

func main() {
	database.StartDB()
	defer database.CloseDB()
	r := router.StartApp()
	r.Run(":8080")
}
