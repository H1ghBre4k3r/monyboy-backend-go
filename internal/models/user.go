package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID            string `gorm:"primaryKey"`
	Username      string `gorm:"unique"`
	DisplayName   string
	Password      string
	Email         string
	EmailVerified bool `gorm:"type:boolean"`
}
