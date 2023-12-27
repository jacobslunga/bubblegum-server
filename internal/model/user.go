package model

import (
  "time"

  "gorm.io/gorm"
  "github.com/google/uuid"
)

type User struct {
  gorm.Model
  ID uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
  Username string
  Email string `gorm:"unique"`
  Password string
  CreatedAt time.Time `gorm:"not null"`
  UpdatedAt time.Time `gorm:"not null"`
  Components []Component `gorm:"foreignkey:CreatorID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  u.ID = uuid.New()
  return
}
