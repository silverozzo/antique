package deck

import "antique/handy"

type MysteryCard struct {
	handy.CardInterface
}

type MysteryDeck struct {
	handy.DeckInterface
}

func NewMysteryDeck() *MysteryDeck {
	deck := MysteryDeck{
		handy.NewDeck([]handy.CardInterface{}),
	}

	return &deck
}
