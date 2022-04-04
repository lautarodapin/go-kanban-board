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
	err := DB.Preload("Board").Preload("Column").First(&ticket, id).Error
	return ticket, err
}

func (t *ticketManager) GetAll() ([]Ticket, error) {
	var tickets []Ticket
	return tickets, DB.Preload("Board").Preload("Column").Find(&tickets).Error
}

func (t *ticketManager) GetAllByQuery(q string) ([]Ticket, error) {
	var tickets []Ticket

	return tickets, DB.
		Preload("Board").
		Joins("Board").
		Find(&tickets).
		Error
}

func (t *ticketManager) DeleteById(id uint64) error {
	return DB.Delete(&Ticket{}, id).Error
}
