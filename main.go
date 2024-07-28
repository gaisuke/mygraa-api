package main

import (
	"github.com/gaisuke/mygram-api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitializeDB()

	r := gin.Default()

	r.Run(":8080")

}
