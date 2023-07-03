package elder

import (
	"antique/elder/deck"
)

type Elder struct {
	mysteries  *deck.MysteryDeck
	researches *deck.ResearchDeck
}

func New() *Elder {
	return &Elder{
		mysteries:  deck.NewMysteryDeck(),
		researches: deck.NewResearchDeck(),
	}
}
