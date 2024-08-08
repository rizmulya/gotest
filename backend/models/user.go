package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Name     string `gorm:"type:varchar(128);not null" json:"name"`
    Email    string `gorm:"type:varchar(64);unique;not null" json:"email"`
    Password string `gorm:"not null" json:"password"`
    Role     string `gorm:"type:varchar(16);not null" json:"role"`
    Image string `json:"image"`
}