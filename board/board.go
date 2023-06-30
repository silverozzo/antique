package board

import (
	"antique/board/deck"
	"antique/handy"
)

// Доска, карта местности (стол)
type Board struct {
	gateStack    *handy.Deck
	clueStack    *handy.Deck
	monsterStack *handy.Deck
}

func New() *Board {
	return &Board{
		gateStack:    deck.NewGateStack(),
		clueStack:    deck.NewClueStack(),
		monsterStack: deck.NewMonsterStack(),
	}
}
