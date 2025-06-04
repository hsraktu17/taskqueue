package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/hsrkatu17/taskqueue/internal/model"
)

// setupRouter initializes a Gin engine with an in-memory SQLite DB for testing.
func setupRouter(t *testing.T) (*gin.Engine, *gorm.DB) {
	gin.SetMode(gin.TestMode)
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open sqlite db: %v", err)
	}
	if err := db.AutoMigrate(&model.Job{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}

	r := gin.New()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.POST("/jobs", CreateJobHandle)
	return r, db
}

func TestCreateJobHandle_Success(t *testing.T) {
	router, db := setupRouter(t)

	payload := map[string]any{"foo": "bar"}
	body, _ := json.Marshal(gin.H{
		"tenant_id": uuid.New(),
		"type":      model.JobTypeSendEmail,
		"payload":   payload,
	})

	req, _ := http.NewRequest(http.MethodPost, "/jobs", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, w.Code)
	}

	// ensure job stored in DB
	var count int64
	if err := db.Model(&model.Job{}).Count(&count).Error; err != nil {
		t.Fatalf("failed to count jobs: %v", err)
	}
	if count != 1 {
		t.Fatalf("expected 1 job in db, got %d", count)
	}
}

func TestCreateJobHandle_BadRequest(t *testing.T) {
	router, _ := setupRouter(t)

	// Missing type and payload fields
	body, _ := json.Marshal(gin.H{"tenant_id": uuid.New()})
	req, _ := http.NewRequest(http.MethodPost, "/jobs", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}
