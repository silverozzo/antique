package board

import (
	"antique/board/deck"
	"antique/handy"
)

// Доска, карта местности (стол)
type Board struct {
	gateStack      *handy.Deck
	clueStack      *handy.Deck
	monsterStack   *handy.Deck
	expeditionDeck *handy.Deck
	doomTrack      int
}

func New(doomTrack int) *Board {
	return &Board{
		gateStack:      deck.NewGateStack(),
		clueStack:      deck.NewClueStack(),
		monsterStack:   deck.NewMonsterStack(),
		expeditionDeck: deck.NewExpeditionDeck(),
		doomTrack:      doomTrack,
	}
}
