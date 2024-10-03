package models

import "time"

type TransferencyModes struct {
    ID         uint      `gorm:"primaryKey"`
    UserID     uint
    User       User      `gorm:"foreignKey:UserID"`
    CPF        string    `gorm:"unique"`
    RandomKey  string    `gorm:"unique"`
    CreatedAt  time.Time
    UpdatedAt  *time.Time
}
