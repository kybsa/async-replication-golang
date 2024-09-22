package entity

type Status int

const (
	CREATED Status = iota + 1
	SENDED
	SYNC
	INVALID_MESSAGE
	ERROR
	MAX_RETRY
)

type SendMessageStatus struct {
	SerialID  uint64 `gorm:"column:serial_id;primaryKey"`
	MessageID uint64 `gorm:"column:message_id"`
	Status    int16  `gorm:"column:status"`
	CreatedAt int64  `gorm:"column:created_at"`
}
