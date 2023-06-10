package database_helper

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"

	_ "github.com/lib/pq"
)

func ConnectionPool(connectionString string, maxConnections, maxLifetime int) (*sqlx.DB, error) {
	// connection pool for Postgres
	connectionPool, err := sqlx.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// setting connection pool configurations
	connectionPool.SetMaxOpenConns(maxConnections)
	connectionPool.SetMaxIdleConns(maxConnections)
	connectionPool.SetConnMaxIdleTime(time.Duration(maxLifetime) * time.Minute)

	return connectionPool, nil
}

func Connection(connectionString string, maxConnections, maxLifetime int, timeout time.Duration) (*sqlx.DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// connection for Postgres
	connection, err := sqlx.ConnectContext(ctx, "postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// setting connection configurations
	connection.SetMaxOpenConns(maxConnections)
	connection.SetMaxIdleConns(maxConnections)
	connection.SetConnMaxIdleTime(time.Duration(maxLifetime) * time.Minute)

	return connection, nil
}
