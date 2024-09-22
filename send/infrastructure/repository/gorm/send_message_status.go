package gorm

import (
	"context"
	"errors"

	appContext "github.com/kybsa/async-replication-golang/send/domain/context"
	"github.com/kybsa/async-replication-golang/send/domain/entity"
	"gorm.io/gorm"
)

type sendMessageStatusImpl struct {
}

func NewSendMessageStatus() *sendMessageStatusImpl {
	return &sendMessageStatusImpl{}
}

func (_self *sendMessageStatusImpl) Add(ctx context.Context, message entity.SendMessageStatus) (entity.SendMessageStatus, error) {
	db, ok := appContext.GetDB(ctx).Instance(ctx).(*gorm.DB)
	if !ok {
		return entity.SendMessageStatus{}, errors.New("invalid db on context")
	}
	db.Save(&message)
	return message, db.Error
}
