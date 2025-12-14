package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"

	"backend/controllers"
	"backend/db"
	"backend/models"
)

func main() {
	db.Connect()
	db.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/users", controllers.CreateUser)
	r.Run(":8080")
}
