package database

import (
	"auth/config"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

func NewPostgres(config *config.Config) *DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.POSTGRES_HOST,
		config.POSTGRES_PORT,
		config.POSTGRES_USER,
		config.POSTGRES_PASS,
		config.POSTGRES_DB_NAME,
		config.POSTGRES_SSLMODE,
	)

	fmt.Println(dsn)
	maxLifetimeConns, err := strconv.Atoi(config.POSTGRES_CONNECTION_MAX_LIFETIME)
	if err != nil {
		panic(err)
	}

	maxOpenConns, err := strconv.Atoi(config.POSTGRES_MAX_OPEN_CONNECTION)
	if err != nil {
		panic(err)
	}
	maxIdleConns, err := strconv.Atoi(config.POSTGRES_MAX_IDLE_CONNECTION)
	if err != nil {
		panic(err)
	}

	db, err := createConnection(dsn,
		maxIdleConns,
		maxOpenConns,
		maxLifetimeConns,
	)

	if err != nil {
		panic(err)
	}
	return &DB{
		DB: db,
	}
}

func createConnection(dsn string, maxIdleConns, maxOpenConns, maxLifetime int) (*sql.DB, error) {

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(time.Second * time.Duration(maxLifetime))

	return db, nil
}
