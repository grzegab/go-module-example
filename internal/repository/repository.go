package repository

import (
	"database/sql"
	"github.com/grzegab/GO_Module_Example/internal/entity"
)

type SchoolRepository interface {
	Connect(*sql.DB) error
	Add(school entity.School) error
	Edit(school entity.School) error
	Suspend(school entity.School) error
}

type ConfigRepository interface {
	Connect(*sql.DB) error
	ChangeTimes(config entity.Config) error
}
