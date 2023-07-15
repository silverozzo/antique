package player

import "antique/board"

// Investigator Сыщик
type Investigator struct {
	isLead       bool
	locationName string
	maxHealth    int
	maxSanity    int
	health       int
	sanity       int
	trainTickets int
	shipTickets  int
}

// ActionType Варианты действий, которые могут делать сыщики
type ActionType int

const (
	TravelAction ActionType = iota
	RestAction
	TradeAction
	PrepareForTravelAction
	AcquireAssetsAction
)

func NewInvestigator(isLead bool) *Investigator {
	return &Investigator{
		isLead:       isLead,
		locationName: "Лондон",
		maxHealth:    5,
		maxSanity:    5,
		health:       5,
		sanity:       5,
	}
}

func (inv *Investigator) BuyTrainTicket(brd *board.Board, prevAct *ActionType) ActionType {
	if prevAct != nil && *prevAct == PrepareForTravelAction {
		panic("нельзя повторять покупку билетов")
	}

	loc := brd.GetLocationByName(inv.locationName)
	if !loc.IsCity() {
		panic("билеты покупать можно только в городе")
	}

	if inv.trainTickets >= 2 {
		panic("Нельзя иметь больше чем 2 билета на поезд")
	}

	inv.trainTickets++

	if inv.shipTickets+inv.trainTickets > 2 {
		inv.shipTickets = 2 - inv.trainTickets
	}

	return PrepareForTravelAction
}

func (inv *Investigator) BuyShipTicket(brd *board.Board, prevAct *ActionType) ActionType {
	if prevAct != nil && *prevAct == PrepareForTravelAction {
		panic("нельзя повторять покупку билетов")
	}

	loc := brd.GetLocationByName(inv.locationName)
	if !loc.IsCity() {
		panic("билеты покупать можно только в городе")
	}

	if inv.shipTickets >= 2 {
		panic("Нельзя иметь больше чем 2 билета на корабль")
	}

	inv.shipTickets++

	if inv.shipTickets+inv.trainTickets > 2 {
		inv.trainTickets = 2 - inv.shipTickets
	}

	return PrepareForTravelAction
}

func (inv *Investigator) Travel(board *board.Board, destName string, prevAct *ActionType) ActionType {
	if prevAct != nil && *prevAct == TravelAction {
		panic("нельзя повторять путешествие")
	}

	moves := board.GetNeighborsOfLocation(inv.locationName)
	for _, item := range moves {
		if item.GetName() == destName {
			inv.locationName = item.GetName()

			return TravelAction
		}
	}

	panic("не найдена соседняя локация по указанному имени: " + destName)
}

func (inv *Investigator) Rest(brd *board.Board, prevAct *ActionType) ActionType {
	if prevAct != nil && *prevAct == RestAction {
		panic("нельзя повторять отдых")
	}

	loc := brd.GetLocationByName(inv.locationName)
	if loc.HasMonsters() {
		panic("нельзя отдыхать в локации с монтрами")
	}

	if inv.health < inv.maxHealth {
		inv.health++
	}

	if inv.sanity < inv.maxSanity {
		inv.sanity++
	}

	return RestAction
}
