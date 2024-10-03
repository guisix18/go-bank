package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    ID                 uint              `gorm:"primaryKey"`
    Name               string
    Email              string            `gorm:"unique"`
    Password           string
    CPF                string
    Age                int
    IsActive           bool              `gorm:"default:false"`
    Cellphone          string            `gorm:"default:''"`
    CreatedAt          time.Time         `gorm:"autoCreateTime"`
    UpdatedAt          *time.Time        `gorm:"autoUpdateTime"`
    DeletedAt          gorm.DeletedAt    `gorm:"index"`
    ImageURL           *string           `gorm:"default:''"`
    CodeVerifications  []CodeVerification `gorm:"foreignKey:UsedBy"`
    Account            *Account
    TransferencyModes  []TransferencyModes
    TransactionsMade   []Transaction      `gorm:"foreignKey:MadeBy;constraint:OnDelete:CASCADE"`
    TransactionsReceived []Transaction    `gorm:"foreignKey:ReceivedBy"`
}