package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
}

func InitDb(config DBConfig) (*sqlx.DB, error) {
	stmt := fmt.Sprintf("host=%s port=%s password=%s dbname=%s user=%s sslmode=disable", config.Host, config.Port, config.Password, config.DBName, config.User)
	db, err := sqlx.Open("postgres", stmt)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
