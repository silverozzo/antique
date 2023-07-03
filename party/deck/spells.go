package deck

import "antique/handy"

type SpellCard struct {
	handy.CardInterface
}

type SpellDeck struct {
	handy.DeckInterface
}

func NewSpellDeck() *SpellDeck {
	deck := SpellDeck{
		handy.NewDeck([]handy.CardInterface{}),
	}

	return &deck
}
