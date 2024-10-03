package models

import "time"

type CodeVerification struct {
    ID          uint      `gorm:"primaryKey"`
    UsedBy      uint      `gorm:"not null"`
    User        User      `gorm:"foreignKey:UsedBy"`
    AlreadyUsed bool      `gorm:"default:false"`
    Code        string    `gorm:"unique"`
    CreatedAt   time.Time
    UsedAt      *time.Time
    ExpireDate  time.Time
    Expired     bool      `gorm:"default:false"`
}
