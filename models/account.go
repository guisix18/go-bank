package models

import (
	"time"
)

type Account struct {
    ID            uint       `gorm:"primaryKey"`
    UserID        uint       `gorm:"unique"`
    User          User       `gorm:"foreignKey:UserID"`
    AccountNumber string     `gorm:"default:cuid()"`
    AccountType   string     `gorm:"type:account_type"`
    Balance       float64    `gorm:"type:numeric"`
    CreatedAt     time.Time  `gorm:"autoCreateTime"`
    UpdatedAt     *time.Time `gorm:"autoUpdateTime"`
}
