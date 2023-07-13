package player

// Investigator Сыщик
type Investigator struct {
	isLead       bool
	locationName string
	health       int
	sanity       int
	trainTickets int
	shipTickets  int
}

func NewInvestigator(isLead bool) *Investigator {
	return &Investigator{
		isLead:       isLead,
		locationName: "Лондон",
		health:       5,
		sanity:       5,
	}
}
func (inv *Investigator) BuyTrainTicket() {
	if inv.trainTickets >= 2 {
		return
	}

	inv.trainTickets++

	if inv.shipTickets+inv.trainTickets > 2 {
		inv.shipTickets = 2 - inv.trainTickets
	}
}

func (inv *Investigator) BuyShipTicket() {
	if inv.shipTickets >= 2 {
		return
	}

	inv.shipTickets++

	if inv.shipTickets+inv.trainTickets > 2 {
		inv.trainTickets = 2 - inv.shipTickets
	}
}
