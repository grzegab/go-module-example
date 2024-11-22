package repository

import (
	"context"
	"database/sql"
	"github.com/grzegab/GO_Module_Example/internal/entity"
	"log"
	"time"
)

type SchoolRepo struct {
	DbDsn   string
	Timeout time.Duration
}

func (s *SchoolRepo) Connect() (*sql.DB, error) {
	log.Println("[School] connecting to database...")
	db, err := sql.Open("pgx", s.DbDsn)
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

func (s *SchoolRepo) NewSchool(school *entity.School, config *entity.Config) (string, error) {
	db, err := s.Connect()
	if err != nil {
		log.Printf("[School] DB error: %v\n", err)
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
	defer cancel()

	var newUUID string

	stmt := `insert into schools (uuid, name, is_active, register_code, admin_id, lesson_length, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7, $8) returning id`

	err = db.QueryRowContext(ctx, stmt,
		school.UUID,
		school.Name,
		school.IsActive,
		school.RegisterCode,
		school.Owner,
		config.LessonLength,
		time.Now(),
		time.Now(),
	).Scan(&newUUID)

	if err != nil {
		return "", err
	}

	return newUUID, nil
}
