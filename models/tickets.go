package models

type ticketManager struct{}

var TicketManager = ticketManager{}

func (t *ticketManager) Create(ticket *Ticket) error {
	return DB.Create(&ticket).Error
}

func (t *ticketManager) Update(ticket *Ticket, id uint64) error {
	return DB.Model(Ticket{}).
		Where("id = ?", id).
		Updates(ticket).
		Error
}

func (t *ticketManager) GetById(id uint64) (Ticket, error) {
	var ticket Ticket
	err := DB.Preload("Kanban").Preload("Dropzone").First(&ticket, id).Error
	return ticket, err
}

func (t *ticketManager) GetAll() ([]Ticket, error) {
	var tickets []Ticket
	return tickets, DB.Preload("Kanban").Preload("Dropzone").Find(&tickets).Error
}

func (t *ticketManager) GetAllByQuery(q string) ([]Ticket, error) {
	var tickets []Ticket
	args := map[string]interface{}{"query": "%" + q + "%"}
	query := `
		title LIKE @query AND description LIKE @query
		AND kanbans.name LIKE @query
	`
	return tickets, DB.
		Preload("Kanban").
		Joins("Kanban").
		Where(query, args).
		Find(&tickets).
		Error
}

func (t *ticketManager) DeleteById(id uint64) error {
	return DB.Delete(&Ticket{}, id).Error
}
