package deck

import "antique/handy"

type ResearchCard struct {
	handy.CardInterface
}

type ResearchDeck struct {
	handy.DeckInterface
}

func NewResearchDeck() *ResearchDeck {
	deck := ResearchDeck{
		handy.NewDeck([]handy.CardInterface{}),
	}

	return &deck
}
