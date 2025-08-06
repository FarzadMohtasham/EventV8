package main

import (
	"github.com/FarzadMohtasham/EventV8/db"
	"github.com/FarzadMohtasham/EventV8/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	serverEngine := gin.Default()

	routes.RegisterRoutes(serverEngine)

	serverEngine.Run(":8880")
}
