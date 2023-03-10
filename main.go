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
	if err := database.DropTable(db, "Booking", "Gedung", "User"); err != nil {
		panic(err)
	}
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to Migrate")
	}
	if err := database.TruncateTableIgnoreFK(db, "gedungs"); err != nil {
		panic(err)
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

	//supClient := supabasestorageuploader.NewSupabaseClient(
	//	"https://ontvftbxgsmzxwlqhsdn.supabase.co",
	//	"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6Im9udHZmdGJ4Z3Ntenh3bHFoc2RuIiwicm9sZSI6ImFub24iLCJpYXQiOjE2Nzg0MDgxMzQsImV4cCI6MTk5Mzk4NDEzNH0.7yypIF1_gkHACVRxolU2KjhLpdUumKw3OdaRtHSnB9Q",
	//	"gambar-gedung",
	//	"",
	//)
	//model.GDummy()
	//model.LDummy()
	v0 := r.Group("/v0")
	v0.POST("/register", Handler.Register)
	v0.POST("/login", Handler.LogIn)
	v0.GET("/validate", middleware.JwtMiddleware(), Handler.Validate)

	v0.GET("/gedungs", Handler.FindAllGedung)
	v0.GET("/gedungs/:id", Handler.GetGedungByID)
	v0.POST("/booking/:id", middleware.JwtMiddleware(), Handler.Booking)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
