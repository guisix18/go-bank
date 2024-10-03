package models

import "time"

type Transaction struct {
    ID                 uint      `gorm:"primaryKey"`
    MadeBy             uint
    SentByUser         User      `gorm:"foreignKey:MadeBy;constraint:OnDelete:CASCADE"`
    ReceivedBy         *uint
    ReceivedByUser     *User     `gorm:"foreignKey:ReceivedBy"`
    CreatedAt          time.Time
    Value              float64   `gorm:"type:numeric"`
    TransactionType    string    `gorm:"type:transaction_type"`
    TransactionStatus  string    `gorm:"type:transaction_status"`
    ScheduleTransaction *time.Time
}
