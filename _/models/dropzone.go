package models

type dropzoneManager struct{}

var DropzoneManager = dropzoneManager{}

func (c *dropzoneManager) Create(dropzone *Dropzone) error {
	return DB.Create(&dropzone).Error
}

func (c *dropzoneManager) Update(dropzone *Dropzone, id string) error {
	return DB.Model(Dropzone{}).
		Where("id = ?", id).
		Updates(dropzone).
		Error
}

func (c *dropzoneManager) GetById(id string) (Dropzone, error) {
	var dropzone Dropzone
	return dropzone, DB.First(&dropzone, id).Error
}

func (c *dropzoneManager) GetAll() ([]Dropzone, error) {
	var dropzones []Dropzone
	return dropzones, DB.Preload("Dropzones").Find(&dropzones).Error
}

func (c *dropzoneManager) DeleteById(id string) error {
	return DB.Delete(&Dropzone{}, id).Error
}
