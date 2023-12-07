package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// constant related
const (
	dbUri                    = "postgresql://chris:password@localhost:5432/mockdb?sslmode=disable"
	defaultMaxConns          = int32(4)
	defaultMinConns          = int32(0)
	defaultMaxConnLifetime   = time.Hour
	defaultMaxConnIdleTime   = time.Minute * 30
	defaultHealthCheckPeriod = time.Minute
	defaultConnectTimeout    = time.Second * 10
)

var conn *pgxpool.Conn

func PoolConfig() *pgxpool.Config {
	dbConfig, err := pgxpool.ParseConfig(dbUri)
	if err != nil {
		log.Fatal("Failed to create a config", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		return true
	}
	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		return true
	}
	dbConfig.BeforeClose = func(c *pgx.Conn) {
		//
	}

	return dbConfig
}
func InitDB() error {
	connPool, err := pgxpool.NewWithConfig(context.Background(), PoolConfig())
	if err != nil {
		log.Fatal("Error while creating connection pool to database", err)
	}

	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from pool ")
	}
	defer connection.Release()

	err = connection.Ping(context.Background())
	if err != nil {
		return err
	}
	conn = connection
	fmt.Println("Database successfully initialized")
	return nil

}

func GetConnection() *pgxpool.Conn {
	return conn

}
