package entity

import "github.com/google/uuid"

type SendMessage struct {
	SerialID       uint64    `gorm:"column:serial_id;primaryKey"`
	IdempotencyKey uuid.UUID `gorm:"type:uuid;column:idempotency_key"`
	ExternalID     string    `gorm:"column:external_id"`
	Message        []byte    `gorm:"column:message"`
	CreatedAt      int64     `gorm:"column:created_at"`
}
