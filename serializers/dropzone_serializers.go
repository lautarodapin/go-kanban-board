package serializers

import "kanban-board/models"

type DropzoneBody struct {
	Name     string `json:"name" binding:"required" validate:"required,min=3,max=255"`
	ColumnID uint   `json:"column_id" binding:"required" validate:"required"`
	Order    uint   `json:"order" binding:"required" validate:"required,min=0"`
}

func (body *DropzoneBody) ToModel() models.Dropzone {
	return models.Dropzone{
		Name:     body.Name,
		ColumnID: body.ColumnID,
		Order:    body.Order,
	}
}
