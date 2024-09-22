package gorm

import (
	"context"
	"errors"

	appContext "github.com/kybsa/async-replication-golang/send/domain/context"

	"github.com/kybsa/async-replication-golang/send/domain/entity"
	"gorm.io/gorm"
)

type sendMessageImpl struct {
}

func NewSendMessage() *sendMessageImpl {
	return &sendMessageImpl{}
}

func (_self *sendMessageImpl) Add(ctx context.Context, message entity.SendMessage) (entity.SendMessage, error) {
	db, ok := appContext.GetDB(ctx).Instance(ctx).(*gorm.DB)
	if !ok {
		return entity.SendMessage{}, errors.New("invalid db on context")
	}
	db.Save(&message)
	return message, db.Error
}

func (_self *sendMessageImpl) ByID(ctx context.Context, ID uint64) (result entity.SendMessage, err error) {
	db, ok := appContext.GetDB(ctx).Instance(ctx).(*gorm.DB)
	if !ok {
		return entity.SendMessage{}, errors.New("invalid db on context")
	}
	err = db.First(&result, ID).Error
	return
}
