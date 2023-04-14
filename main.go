package main

import (
	"MygarmProject/database"
	"MygarmProject/router"
	"os"
)

func main() {
	database.StartDB()
	defer database.CloseDB()
	r := router.StartApp()
	var PORT = os.Getenv("PORT")
	r.Run(":" + PORT)
}
