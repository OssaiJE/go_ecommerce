package main

import (
	"go_ecommerce/config"
	"go_ecommerce/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
    config.LoadEnv()
	config.ConnectDB()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "2010"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.Routes(router)

	router.Run(":" + port)
}
