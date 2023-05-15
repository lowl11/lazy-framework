package postgres_service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/lazy-framework/log"
	"time"

	_ "github.com/lib/pq"
)

func (service *Service) ConnectionPool() (*sqlx.DB, error) {
	// connection pool for Postgres
	connectionPool, err := sqlx.Open("postgres", service.connectionString)
	if err != nil {
		return nil, err
	}

	// setting connection pool configurations
	connectionPool.SetMaxOpenConns(service.maxConnections)
	connectionPool.SetMaxIdleConns(service.maxConnections)
	connectionPool.SetConnMaxIdleTime(time.Duration(service.maxLifetime) * time.Minute)

	// ping database
	log.Info("Ping Postgres database connection pool...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err = connectionPool.PingContext(ctx); err != nil {
		return nil, err
	}
	log.Info("Ping Postgres database connection pool done!")

	return connectionPool, nil
}

func (service *Service) Connection() (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// connection for Postgres
	connection, err := sqlx.ConnectContext(ctx, "postgres", service.connectionString)
	if err != nil {
		return nil, err
	}

	// setting connection configurations
	connection.SetMaxOpenConns(service.maxConnections)
	connection.SetMaxIdleConns(service.maxConnections)
	connection.SetConnMaxIdleTime(time.Duration(service.maxLifetime) * time.Minute)

	log.Info("Ping Postgres database connection...")
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err = connection.PingContext(ctx); err != nil {
		return nil, err
	}
	log.Info("Ping Postgres database connection done!")

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
