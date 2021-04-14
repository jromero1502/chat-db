package db

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Lastname  string
	Email     string
	Birthday  time.Time
	LastLogin time.Time
}
