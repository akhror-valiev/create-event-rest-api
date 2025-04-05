package main

import (
	"github.com/akhror-valiev/rest-api/db"
	"github.com/akhror-valiev/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") // localhost:8080
}



