package main

import (
	"InternBCC/Handler"
	"InternBCC/database"
	"InternBCC/middleware"
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

	//model.GDummy()
	//model.LDummy()
	v1 := r.Group("/v1")
	v1.POST("/register", Handler.Register)
	v1.POST("/login", Handler.LogIn)
	v1.GET("/validate", middleware.JwtMiddleware(), Handler.Validate)
	v1.PUT("/update", middleware.JwtMiddleware(), Handler.ChangeNameNumber)
	v1.PUT("/change-pass", middleware.JwtMiddleware(), Handler.ChangePass)
	v1.POST("/delete-account", middleware.JwtMiddleware(), Handler.DeleteAccount)

	v1.GET("/gedungs", Handler.FindAllGedung)
	v1.GET("/gedungs/:id", Handler.GetGedungByID)
	v1.GET("/history", middleware.JwtMiddleware(), Handler.GetHistory)
	v1.POST("/booking/:id", middleware.JwtMiddleware(), Handler.Booking)
	v1.GET("/booking-details", middleware.JwtMiddleware(), Handler.GetBookingData)
	v1.POST("/payment/:id", middleware.JwtMiddleware(), Handler.Payment)
	r.Run()
}
