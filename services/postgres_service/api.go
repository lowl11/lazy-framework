package postgres_service

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-framework/log"
	"time"

	_ "github.com/lib/pq"
)

func (service *Service) Connect() (*sqlx.DB, error) {
	// подключение к Postgres
	connection, err := sqlx.Open("postgres", service.connectionString)
	if err != nil {
		return nil, err
	}

	connection.SetMaxOpenConns(service.maxConnections)
	connection.SetMaxIdleConns(service.maxConnections)
	connection.SetConnMaxIdleTime(time.Duration(service.maxLifetime) * time.Minute)

	log.Info("Ping Postgres database...")
	if err = connection.Ping(); err != nil {
		return nil, err
	}
	log.Info("Ping Postgres database done!")

	return connection, nil
}

func (service *Service) SetMaxConnections(maxConnections int) *Service {
	service.maxConnections = maxConnections
	return service
}

func (service *Service) SetMaxLifetime(maxLifetime int) *Service {
	service.maxLifetime = maxLifetime
	return service
}
