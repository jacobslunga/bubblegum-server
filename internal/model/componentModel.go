package model

import (
  "time"

  "github.com/google/uuid"
  "gorm.io/gorm"
)

type CreateComponentRequest struct {
  Name        string `json:"name"`
  Description string `json:"description"`
  Code        string `json:"code"`
}

type Component struct {
  gorm.Model
  ID          string    `gorm:"type:char(36);primary_key"`
  Name        string    `gorm:"not null"`
  CreatorID   uuid.UUID `gorm:"not null"`
  CreatedAt   time.Time `gorm:"not null"`
  UpdatedAt   time.Time `gorm:"not null"`
  Description string
  Code        string `gorm:"not null"`
  Creator     User   `gorm:"foreignkey:CreatorID"`
}

func (c *Component) BeforeCreate(tx *gorm.DB) (err error) {
  c.ID = uuid.New().String()
  return
}
