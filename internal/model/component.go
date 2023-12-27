package model

import (
  "time"

  "gorm.io/gorm"
  "github.com/google/uuid"
)

type Component struct {
  gorm.Model
  ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
  Name string `gorm:"not null"`
  CreatorID uuid.UUID `gorm:"not null"`
  CreatedAt time.Time `gorm:"not null"`
  UpdatedAt time.Time `gorm:"not null"`
  Description string
  Code string `gorm:"not null"`
  Creator User `gorm:"foreignkey:CreatorID"`
}

func (c *Component) BeforeCreate(tx *gorm.DB) (err error) {
  c.ID = uuid.New()
  return
}
