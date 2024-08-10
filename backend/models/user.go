package models

import "gorm.io/gorm"

type User struct {
    gorm.Model // id,createdAt,updatedAt,deletedAt
    Uid string `gorm:"type:varchar(32);uniqueIndex;not null" json:"uid"` // use `uid` (random string) identifier instead of integer `id` for security purposes
    Name     string `gorm:"type:varchar(128);not null" json:"name"`
    Email    string `gorm:"type:varchar(64);unique;not null" json:"email"`
    Password string `gorm:"not null" json:"password"`
    Role     string `gorm:"type:varchar(16);not null" json:"role"`
    Image string `json:"image"`
}