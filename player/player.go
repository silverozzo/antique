package player

import "antique/board"

// Player Игрок... или что-то за него
type Player struct {
	investigator *Investigator
}

func NewPlayer(investigator *Investigator) *Player {
	return &Player{
		investigator: investigator,
	}
}

var step = 0

func (p *Player) Action(brd *board.Board, prevAct *ActionType) ActionType {
	var act ActionType

	switch step {
	case 0:
		act = p.investigator.BuyShipTicket(brd, prevAct)

	case 1:
		act = p.investigator.Travel(brd, "Глушь", prevAct)

	case 2:
		act = p.investigator.Rest(brd, prevAct)

	default:
		panic("игрок не значет что делать в фазу действий")
	}

	step++

	return act
}
