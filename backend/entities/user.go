package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	FirstName *string `gorm:"not null" `
	LastName  *string `gorm:"not null" `
	Email     *string `gorm:"unique;not null`
	Password  *string `gorm:"not null" `
	RoleId    *uint   `gorm:"not null" `
	Active    bool   `gorm:"not null;default:false"`
}
