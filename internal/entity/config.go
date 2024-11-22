package entity

import (
	"time"
)

type Config struct {
	StartMinute  int       `json:"start_minute"`
	LessonLength int       `json:"lesson_length"`
	CreatedAt    time.Time `json:"-"`
	DeletedAt    time.Time `json:"-"`
}

func NewBasicConfig() *Config {
	return &Config{
		StartMinute:  0,
		LessonLength: 60,
		CreatedAt:    time.Time{},
		DeletedAt:    time.Time{},
	}
}
