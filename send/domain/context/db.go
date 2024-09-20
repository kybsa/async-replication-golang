package context

import (
	"context"

	"github.com/kybsa/async-replication-golang/send/domain/repository"
)

type ctxKey string

const keyDB = ctxKey("_DbKey_")

func AddDB(ctx context.Context, database repository.Database) context.Context {
	return context.WithValue(ctx, keyDB, database)
}

func GetDB(ctx context.Context) repository.Database {
	anyDB := ctx.Value(keyDB)
	if anyDB == nil {
		return nil
	}
	db, ok := anyDB.(repository.Database)

	if !ok {
		return nil
	}
	return db
}
