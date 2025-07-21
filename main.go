package main

import (
	"chatbot/src/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	routes.SetupRoutes(router)

	err := router.Run(":5001")
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
