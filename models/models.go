package models

import "time"

type Kanban struct {
	BaseModel
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	StartDate time.Time `json:"start_date" gorm:"not null"`
	EndDate   time.Time `json:"end_date" gorm:"not null"`
}

type Column struct {
	BaseModel
	Name      string     `json:"name" gorm:"type:varchar(255);not null"`
	Order     uint       `json:"order" gorm:"type:integer;not null;default:0;unique;"`
	Dropzones []Dropzone `json:"dropzones" gorm:"foreignkey:ColumnID"`
}

type Dropzone struct {
	BaseModel
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	ColumnID uint   `json:"column_id" gorm:"type:integer;not null"`
	Order    uint   `json:"order" gorm:"type:integer;not null;default:0;"`
	Column   Column `json:"column" gorm:"foreignkey:ColumnID"`
}

type Ticket struct {
	BaseModel
	Title       string   `json:"title" gorm:"type:varchar(255);not null"`
	Description string   `json:"description" gorm:"type:varchar(1025)"`
	KanbanID    uint     `json:"kanban_id" gorm:"type:integer"`
	DropzoneID  uint     `json:"dropzone_id" gorm:"type:integer; not null"`
	Kanban      Kanban   `json:"kanban" gorm:"foreignkey:KanbanID"`
	Dropzone    Dropzone `json:"dropzone" gorm:"foreignkey:DropzoneID"`
}
