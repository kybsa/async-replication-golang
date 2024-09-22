package gorm

import (
	"context"

	gormio "gorm.io/gorm"
)

type databaseImpl struct {
	db *gormio.DB
}

func NewDatabase(db *gormio.DB) *databaseImpl {
	return &databaseImpl{
		db: db,
	}
}

func (_self *databaseImpl) Begin(ctx context.Context) error {
	return _self.db.Begin().Error
}

func (_self *databaseImpl) Commit(ctx context.Context) (err error) {
	return _self.db.Commit().Error
}

func (_self *databaseImpl) Rollback(ctx context.Context) (err error) {
	return _self.db.Rollback().Error
}

func (_self *databaseImpl) Instance(ctx context.Context) interface{} {
	return _self.db
}
