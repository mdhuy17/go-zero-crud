package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string         `gorm:"primary_key;size:100"`
	CreatedAt *time.Time     `json:"created_at" json:"created_at"`
	UpdatedAt *time.Time     `json:"updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" yaml:"deleted_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.ID = uuid.New().String()
	return nil
}
