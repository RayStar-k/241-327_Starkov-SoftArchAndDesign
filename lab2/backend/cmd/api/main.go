package main

import (
	"log"

	"guitarshop/internal/config"
	"guitarshop/internal/database"
	"guitarshop/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := database.Connect(&cfg.Database); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	gin.SetMode(cfg.Server.Mode)

	router := gin.Default()

	// Корневой маршрут
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Guitar Shop API",
			"version": "1.0.0",
		})
	})
	
	api := router.Group("/api")
	{
		guitars := api.Group("/guitars")
		{
			guitars.GET("", handlers.GetGuitars)
			guitars.GET("/:id", handlers.GetGuitar)
			guitars.POST("", handlers.CreateGuitar)
			guitars.PUT("/:id", handlers.UpdateGuitar)
			guitars.DELETE("/:id", handlers.DeleteGuitar)
		}
	}

	serverAddr := ":" + cfg.Server.Port
	log.Printf("Starting server on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
