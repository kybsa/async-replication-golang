package repository

import (
	"context"

	"github.com/kybsa/async-replication-golang/send/domain/entity"
)

type SendMessage interface {
	Add(ctx context.Context, message entity.SendMessage) (entity.SendMessage, error)
	ByID(ctx context.Context, ID uint64) (entity.SendMessage, error)
}
