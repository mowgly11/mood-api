package main

import (
	//"fmt"
	"project/database"
	"project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	
	engine := gin.Default()

	routes.RegisterRoutes(engine)

	engine.Run(":8000")
}