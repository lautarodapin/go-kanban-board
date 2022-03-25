package models

import "time"

type CreateKanban struct {
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	StartDate time.Time `json:"start_date" gorm:"not null"`
	EndDate   time.Time `json:"end_date" gorm:"not null"`
}

type UpdateKanban struct {
	Name    string    `json:"name" gorm:"type:varchar(255)"`
	EndDate time.Time `json:"end_date" gorm:"not null"`
}

type Kanban struct {
	BaseModel
	CreateKanban
}

type BaseColumn struct {
	Name  string `json:"name" gorm:"type:varchar(255)"`
	Order uint   `json:"order" gorm:"type:integer;not null;default:0;unique;"`
}

type Column struct {
	BaseModel
	BaseColumn
}

type BaseDropzone struct {
	Name     string `json:"name" gorm:"type:varchar(255)"`
	ColumnID uint   `json:"column_id" gorm:"type:integer"`
	Order    uint   `json:"order" gorm:"type:integer;not null;default:0;"`
}

type Dropzone struct {
	BaseModel
	BaseDropzone
	Column Column `json:"column" gorm:"foreignkey:ColumnID"`
}

type BaseTicket struct {
	Title       string `json:"title" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:varchar(1025)"`
	KanbanID    uint   `json:"kanban_id" gorm:"type:integer"`
	DropzoneID  uint   `json:"dropzone_id" gorm:"type:integer"`
}

type Ticket struct {
	BaseModel
	BaseTicket
	Kanban   Kanban   `json:"kanban" gorm:"foreignkey:KanbanID"`
	Dropzone Dropzone `json:"dropzone" gorm:"foreignkey:DropzoneID"`
}
