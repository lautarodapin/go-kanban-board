package models

type boardManager struct{}

var BoardManager = boardManager{}

func (k *boardManager) Create(board *Board) error {
	return DB.Create(&board).Error
}

func (k *boardManager) Update(board *Board, id uint) error {
	return DB.Model(Board{}).
		Where("id = ?", id).
		Updates(board).
		Error
}

func (k *boardManager) GetById(id uint) (Board, error) {
	var board Board
	return board, DB.Preload("Columns").First(&board, id).Error
}

func (k *boardManager) GetAll() ([]Board, error) {
	var boards []Board
	return boards, DB.Preload("Columns").Find(&boards).Error
}

func (k *boardManager) DeleteById(id uint) error {
	return DB.Delete(&Board{}, id).Error
}
