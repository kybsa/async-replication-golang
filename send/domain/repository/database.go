package repository

import "context"

type Database interface {
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	Instance(ctx context.Context) interface{}
}
