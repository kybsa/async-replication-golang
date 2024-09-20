package repository

import (
	"context"

	"github.com/kybsa/async-replication-golang/send/domain/entity"
)

type SendMessageStatus interface {
	Add(ctx context.Context, message entity.SendMessageStatus) (entity.SendMessageStatus, error)
}
