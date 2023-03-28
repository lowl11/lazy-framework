package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-framework/events/script_event"
	"github.com/lowl11/lazy-framework/log"
	"strings"
	"time"
)

type Base struct {
	script *script_event.Event
}

func CreateBase(script *script_event.Event) Base {
	return Base{
		script: script,
	}
}

func (repo *Base) StartScript(name string) string {
	return repo.script.StartScript(name)
}

func (repo *Base) Script(folder, name string) string {
	return repo.script.Script(folder, name)
}

func (repo *Base) Ctx(customTimeout ...time.Duration) (context.Context, func()) {
	defaultTimeout := time.Second * 5
	if len(customTimeout) > 0 {
		defaultTimeout = customTimeout[0]
	}
	return context.WithTimeout(context.Background(), defaultTimeout)
}

func (repo *Base) CloseRows(rows *sqlx.Rows) {
	if err := rows.Close(); err != nil {
		log.Error(err, "Closing rows error")
	}
}

func (repo *Base) Rollback(transaction *sqlx.Tx) {
	if err := transaction.Rollback(); err != nil {
		if !strings.Contains(err.Error(), "sql: transaction has already been committed or rolled back") {
			log.Error(err, "Rollback transaction error")
		}
	}
}

func (repo *Base) Transaction(connection *sqlx.DB, transactionActions func(tx *sqlx.Tx) error) error {
	transaction, err := connection.Beginx()
	if err != nil {
		return err
	}
	defer repo.Rollback(transaction)

	if err = transactionActions(transaction); err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}
