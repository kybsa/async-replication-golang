package repository

import (
	"context"

	"github.com/kybsa/async-replication-golang/send/domain/entity"
)

type Queue interface {
	Send(ctx context.Context, message entity.SendMessage) error
}
