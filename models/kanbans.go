package models

type kanbanManager struct{}

var KanbanManager = kanbanManager{}

func (k *kanbanManager) Create(kanban *Kanban) error {
	return DB.Create(&kanban).Error
}

func (k *kanbanManager) Update(kanban *Kanban, id uint) error {
	return DB.Model(Kanban{}).
		Where("id = ?", id).
		Updates(kanban).
		Error
}

func (k *kanbanManager) GetById(id uint) (Kanban, error) {
	var kanban Kanban
	return kanban, DB.Preload("Columns").First(&kanban, id).Error
}

func (k *kanbanManager) GetAll() ([]Kanban, error) {
	var kanbans []Kanban
	return kanbans, DB.Preload("Columns").Find(&kanbans).Error
}

func (k *kanbanManager) DeleteById(id uint) error {
	return DB.Delete(&Kanban{}, id).Error
}
