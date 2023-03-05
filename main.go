package main

import (
	"InternBCC/Handler"
	"InternBCC/database"
	"InternBCC/middleware"
	"InternBCC/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	db := database.InitDB()
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to Migrate")
	}
	if err != nil {
		log.Fatalln("failed to load env file")
	}
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.AbortWithStatus(204)
		} else {
			c.Next()
		}
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	model.TagDummy()
	model.GDummy()
	v0 := r.Group("/v0")
	v0.POST("/register", Handler.Register)
	v0.POST("/login", Handler.LogIn)
	v0.GET("/logout", Handler.LogOut)
	v0.GET("/validate", middleware.JwtMiddleware(), Handler.Validate)
	//v1.POST("/changePass", Handler.ChangePass)
	//asd
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
