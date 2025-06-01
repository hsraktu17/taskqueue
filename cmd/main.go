package main

import (
	"log"

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

	if err := db.AutoMigrate(&model.Job{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
	r.Run(":" + cfg.Port)
}
