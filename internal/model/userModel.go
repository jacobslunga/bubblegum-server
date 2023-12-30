package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	ID         string      `gorm:"type:char(36);primary_key" json:"id"`
	Username   string      `gorm:"unique" json:"username"`
	Email      string      `gorm:"unique" json:"email"`
	Password   string      `gorm:"not null" json:"password"`
	CreatedAt  time.Time   `gorm:"not null" json:"created_at"`
	UpdatedAt  time.Time   `gorm:"not null" json:"updated_at"`
	Components []Component `gorm:"foreignkey:CreatorID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New().String()
	return
}

func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}

	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
