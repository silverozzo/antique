package elder

import (
	"antique/elder/deck"
	"antique/handy"
)

type Elder struct {
	mysteries  *handy.Deck
	researches *handy.Deck
	specials   map[string]*handy.Deck
}

func New() *Elder {
	return &Elder{
		mysteries:  deck.NewMysteryDeck(),
		researches: deck.NewResearchDeck(),
	}
}
