package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Job struct {
	ID               uuid.UUID      `gorm:"type:uuid;primaryKey"`
	TenantID         uuid.UUID      `gorm:"type:uuid;not null;index:idx_jobs_tenant_status_visible,priority:1"`
	Type             JobType        `gorm:"type:varchar(32);not null"`
	Payload          datatypes.JSON `gorm:"type:jsonb;not null"`
	Status           JobStatus      `gorm:"type:varchar(16);not null;default:'pending';index:idx_jobs_tenant_status_visible,priority:2"`
	DeduplicationKey *string        `gorm:"type:varchar(128);uniqueIndex:idx_jobs_dedup,where:deduplication_key IS NOT NULL"`
	VisibleAt        time.Time      `gorm:"type:timestamp;not null;default:now();index:idx_jobs_tenant_status_visible,priority:3"`
	RetryCount       int            `gorm:"type:int;not null;default:0"`
	LastError        *string        `gorm:"type:text"`
	CreatedAt        time.Time      `gorm:"type:timestamp;not null;default:now()"`
	UpdatedAt        time.Time      `gorm:"type:timestamp;not null;default:now()"`
}
