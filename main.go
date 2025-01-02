package main

import (
	"log"

	"github.com/1strewave/AuthLite/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	router.POST("/register", routes.Register)
	router.POST("/login", routes.Login)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
