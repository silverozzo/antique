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

func (pl *Player) Travel(board *board.Board, destName string) {
	moves := board.GetNeighborsOfLocation(pl.investigator.locationName)
	for _, item := range moves {
		if item.GetName() == destName {
			pl.investigator.locationName = item.GetName()
		}
	}

	panic("не найдена соседняя локация по указанному имени: " + destName)
}
