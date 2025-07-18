package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/hsrkatu17/taskqueue/internal/model"
)

// CreateJobRequest represents the expected payload for creating a job.
type CreateJobRequest struct {
	TenantID uuid.UUID      `json:"tenant_id" binding:"required"`
	Type     model.JobType  `json:"type" binding:"required"`
	Payload  datatypes.JSON `json:"payload" binding:"required"`
}

// CreateJobHandle stores a new job in the database.
func CreateJobHandle(c *gin.Context) {
	var req CreateJobRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbValue, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "database key missing in context"})
		return
	}
	db, ok := dbValue.(*gorm.DB)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid database object in context"})
		return
	}

	job := model.Job{
		TenantID:  req.TenantID,
		Type:      req.Type,
		Payload:   req.Payload,
		Status:    model.StatusPending,
		VisibleAt: time.Now(),
	}

	if err := db.Create(&job).Error; err != nil {
		log.Printf("Failed to create job: %v", err) // Log the detailed error internally
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"}) // Return a generic error message to the client
		return
	}

	c.JSON(http.StatusCreated, job)
}
