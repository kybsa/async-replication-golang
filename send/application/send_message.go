package application

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	appContext "github.com/kybsa/async-replication-golang/send/domain/context"
	"github.com/kybsa/async-replication-golang/send/domain/dto"

	"github.com/kybsa/async-replication-golang/send/domain/entity"
	"github.com/kybsa/async-replication-golang/send/domain/repository"
)

type SendMessage interface {
}

type sendMessageImpl struct {
	sendMessageRepository       repository.SendMessage
	sendMessageStatusRepository repository.SendMessageStatus
	queue                       repository.Queue
	database                    repository.Database
}

func NewMessage(sendMessageRepository repository.SendMessage, sendMessageStatusRepository repository.SendMessageStatus, queue repository.Queue, database repository.Database) *sendMessageImpl {
	return &sendMessageImpl{
		sendMessageRepository:       sendMessageRepository,
		sendMessageStatusRepository: sendMessageStatusRepository,
		queue:                       queue,
		database:                    database,
	}
}

func (_self *sendMessageImpl) CreateMessage(ctx context.Context, message dto.Message, db repository.Database) (entityMessage entity.SendMessage, err error) {

	uuidSendMessage, err := uuid.NewV7()
	if err != nil {
		return
	}

	now := time.Now()
	entityMessage.IdempotencyKey = uuidSendMessage
	entityMessage.CreatedAt = now.Unix()
	entityMessage.Message = message.Content
	entityMessage.ExternalID = message.ExternalID

	ctx = appContext.AddDB(ctx, db)

	entityMessage, err = _self.sendMessageRepository.Add(ctx, entityMessage)
	if err != nil {
		return
	}

	err = _self.addStatus(ctx, now.Unix(), entity.CREATED, entityMessage)
	if err != nil {
		return
	}

	return entityMessage, nil
}

func (_self *sendMessageImpl) SendMessage(ctx context.Context, message entity.SendMessage) (err error) {
	ctx = appContext.AddDB(ctx, _self.database)
	entityMessage, err := _self.sendMessageRepository.ByID(ctx, message.SerialID)
	if err != nil {
		return
	}

	err = _self.database.Begin(ctx)
	if err != nil {
		return
	}

	if err = _self.queue.Send(ctx, entityMessage); err == nil {
		if err = _self.addStatus(ctx, time.Now().Unix(), entity.SENDED, entityMessage); err == nil {
			err = _self.database.Commit(ctx)
		} else {
			errRollBack := _self.database.Rollback(ctx)
			err = errors.Join(err, errRollBack)
		}
	}
	return err
}

func (_self *sendMessageImpl) addStatus(ctx context.Context, createdAt int64, status entity.Status, entityMessage entity.SendMessage) (err error) {
	entityMessageStatus := entity.SendMessageStatus{
		MessageID: entityMessage.SerialID,
		Status:    int16(status),
		CreatedAt: createdAt,
	}

	_, err = _self.sendMessageStatusRepository.Add(ctx, entityMessageStatus)
	return
}
