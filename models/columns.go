package models

type columnManager struct{}

var ColumnManager = columnManager{}

func (c *columnManager) Create(column *Column) error {
	return DB.Create(&column).Error
}

func (c *columnManager) Update(column *Column, id string) error {
	return DB.Model(Column{}).
		Where("id = ?", id).
		Updates(column).
		Error
}

func (c *columnManager) GetById(id string) (Column, error) {
	var column Column
	err := DB.First(&column, id).Error
	return column, err
}

func (c *columnManager) GetAll() ([]Column, error) {
	var columns []Column
	return columns, DB.Preload("Dropzones").Find(&columns).Error
}

func (c *columnManager) DeleteById(id string) error {
	return DB.Delete(&Column{}, id).Error
}
