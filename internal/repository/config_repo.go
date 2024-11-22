package repository

import (
	"database/sql"
	"log"
	"time"
)

type ConfigRepo struct {
	dbDsn   string
	timeout time.Duration
}

func (cr *ConfigRepo) Connect() (*sql.DB, error) {
	log.Println("[School] connecting to database...")
	db, err := sql.Open("pgx", cr.dbDsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("[School] Connected to DB!")

	return db, nil
}
