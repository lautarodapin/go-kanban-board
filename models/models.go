package models

import "time"

type Board struct {
	BaseModel
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	StartDate time.Time `json:"start_date" gorm:"not null"`
	EndDate   time.Time `json:"end_date" gorm:"not null"`
}

type Column struct {
	BaseModel
	Name  string `json:"name" gorm:"type:varchar(255);not null"`
	Order uint   `json:"order" gorm:"type:integer;not null;default:0;unique;"`
}

type Ticket struct {
	BaseModel
	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	Description string `json:"description" gorm:"type:varchar(1025)"`
	BoardID     uint   `json:"board_id" gorm:"type:integer"`
	DropzoneID  uint   `json:"dropzone_id" gorm:"type:integer; not null"`
	Board       Board  `json:"board" gorm:"foreignkey:BoardID"`
}
