package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/lowl11/lazy-framework/log"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type NoSqlBase struct {
	//
}

func NewNoSqlBase() NoSqlBase {
	return NoSqlBase{}
}

func (repo *NoSqlBase) Guid() string {
	return uuid.New().String()
}

func (repo *NoSqlBase) Ctx(customTimeout ...time.Duration) (context.Context, func()) {
	defaultTimeout := time.Second * 5
	if len(customTimeout) > 0 {
		defaultTimeout = customTimeout[0]
	}
	return context.WithTimeout(context.Background(), defaultTimeout)
}

func (repo *NoSqlBase) CloseCursor(cursor *mongo.Cursor) {
	ctx, cancel := repo.Ctx(time.Second * 2)
	defer cancel()

	if err := cursor.Close(ctx); err != nil {
		log.Error(err, "Close cursor error")
	}
}

func (repo *NoSqlBase) LogError(err error) {
	if err != nil {
		log.Error(err, "Cursor error")
	}
}
