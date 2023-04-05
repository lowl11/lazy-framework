package interfaces

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type IRepository interface {
	Ctx(customTimeout ...time.Duration) (context.Context, func())
}

type ISqlRepository interface {
	StartScript(name string) string
	Script(folder, name string) string

	CloseRows(rows *sqlx.Rows)
	Rollback(transaction *sqlx.Tx)
	Transaction(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error
}
