package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UsingMiddlewares() {
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "CORS works!"})
	})
	router.Run(":5000")
}

// Custom Middleware
func FindUserAgent() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println(ctx.GetHeader("User-Agent"))
		ctx.Next()
	}
}

func UsingCustomMiddlewares() {
	router := gin.Default()
	router.Use(FindUserAgent())
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Custom middleware works!"})
	})
	router.Run(":5000")
}
