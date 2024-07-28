package main

import (
	"github.com/gaisuke/mygram-api/database"
	"github.com/gaisuke/mygram-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitializeDB()

	r := gin.Default()

	routes.SetupUserRoutes(r)

	r.Run(":8080")

}
