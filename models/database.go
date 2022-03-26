package models

import (
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DATABASE_URL")
	var database *gorm.DB
	var err error
	if dsn == "" {
		database, err = gorm.Open(sqlite.Open("gorm.sqlite"), &gorm.Config{})
	} else {
		database, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gorm.Config{})
	}

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Kanban{}, &Column{}, &Dropzone{}, &Ticket{})

	DB = database
}
