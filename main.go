package main

import (
	"os"

	"github.com/Caknoooo/nasabah-bank/config"
	"github.com/Caknoooo/nasabah-bank/middlewares"
	"github.com/Caknoooo/nasabah-bank/router"
)

func main(){
	database := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(database)
	
	server := router.Routes(database)
	
	server.Use(middlewares.CORSMiddleware())
	
	port := os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	server.Run(":" + port)
}