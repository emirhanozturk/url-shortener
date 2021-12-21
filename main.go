package main

import (
	"fmt"
	handlers "url-shortener/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to url-shortener API",
		})
	})

	r.POST("/create-short-url", func(c *gin.Context) {
		handlers.CreateShortUrl(c)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handlers.RedirectMainUrl(c)
	})

	err := r.Run(":8080")
	if err != nil {
		panic(fmt.Sprintf("Failed to start web server - Error %v", err))
	}

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
