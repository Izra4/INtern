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
	if err := database.DropTable(db, "Booking", "Gedung", "User"); err != nil {
		panic(err)
	}
	if err := database.TruncateTableIgnoreFK(db, "gedungs"); err != nil {
		panic(err)
	}
	if err := database.TruncateTableIgnoreFK(db, "links"); err != nil {
		panic(err)
	}
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

	model.GDummy()
	model.LDummy()
	v0 := r.Group("/v0")
	v0.POST("/register", Handler.Register)
	v0.POST("/login", Handler.LogIn)
	v0.GET("/validate", middleware.JwtMiddleware(), Handler.Validate)
	v0.PATCH("/update", middleware.JwtMiddleware(), Handler.ChangeNameNumber)

	v0.GET("/gedungs", Handler.FindAllGedung)
	v0.GET("/gedungs/:id", Handler.GetGedungByID)
	v0.GET("/history", middleware.JwtMiddleware(), Handler.GetHistory)
	v0.POST("/booking/:id", middleware.JwtMiddleware(), Handler.Booking)
	v0.POST("/payment/:id", middleware.JwtMiddleware(), Handler.Payment)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
