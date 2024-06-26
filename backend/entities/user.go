package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint    `gorm:"primaryKey"`
	FirstName *string `gorm:"not null" `
	LastName  *string `gorm:"not null" `
	Email     *string `gorm:"unique;not null`
	Password  *string `gorm:"not null" `
	RoleId    *uint   `gorm:"not null" `
	Active    bool    `gorm:"not null;default:false"`
}

type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FullName  string    `json:"fullName"`
	RoleId    uint      `json:"roleId"`
	CreatedAt time.Time `json:"createdAt"`
	Active    bool      `json:"active"`
}
