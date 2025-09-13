package main

import (
	"azizur/rest-api/db"
	"azizur/rest-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	engine := gin.Default()
	routes.RegisterRoutes(engine)
	engine.Run(":8080") // localhost:8080
}
