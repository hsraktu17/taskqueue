package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hsrkatu17/taskqueue/internal/config"
)

func main() {
	r := gin.Default()
	cfg := config.Load()
	log.Println("config", cfg)
	r.Run(":" + cfg.Port)
}
