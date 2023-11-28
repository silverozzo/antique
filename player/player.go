package player

import (
	"antique/board"
)

// Player Игрок... или что-то за него
type Player struct {
	investigator *Investigator
}

func NewPlayer(investigator *Investigator) *Player {
	return &Player{
		investigator: investigator,
	}
}

var commands = []Command{
	{
		action: TravelAction,
		arguments: map[string]interface{}{
			"destination": "Глушь",
		},
	},
	{
		action: RestAction,
	},
	{
		action: TravelAction,
		arguments: map[string]interface{}{
			"destination": "Лондон",
		},
	},
	{
		action: FocusAction,
	},
	{
		action: TravelAction,
		arguments: map[string]interface{}{
			"destination": "Море",
		},
	},
	{
		action: RestAction,
	},
}

var step = 0

func (p *Player) Action(brd *board.Board, prevAct *ActionType) ActionType {
	var act ActionType

	command := commands[step]
	step++

	if step >= len(commands) {
		panic("не осталось возможных ходов для игрока")
	}

	switch command.action {
	case RestAction:
		act = p.investigator.Rest(brd, prevAct)

	case FocusAction:
		act = p.investigator.Focus(brd, prevAct)

	case PrepareForTravelAction:
		if command.arguments["ticker"] == "train" {
			act = p.investigator.BuyTrainTicket(brd, prevAct)
		} else if command.arguments["ticket"] == "ship" {
			act = p.investigator.BuyShipTicket(brd, prevAct)
		} else {
			panic("не определен тип билета для подготовки к путешествию")
		}

	case TravelAction:
		act = p.investigator.Travel(brd, command.arguments["destination"].(string), prevAct)

	default:
		panic("игрок не знает как выполнить команду")
	}

	return act
}
