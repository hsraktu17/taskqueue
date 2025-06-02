package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsrkatu17/taskqueue/internal/config"
	"github.com/hsrkatu17/taskqueue/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	cfg := config.Load()
	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to db %v", err)
	}

	// REMOVE this block if enums/tables are already created via SQL migration:
	if err := db.AutoMigrate(&model.Job{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	log.Println("Database connected.")

	// Attach DB to Gin context if you want (recommended for API handlers)
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":" + cfg.Port)
}
