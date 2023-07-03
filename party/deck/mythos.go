package deck

import "antique/handy"

type MythCard struct {
	handy.CardInterface
}

type MythDeck struct {
	handy.DeckInterface
}

func NewMythDeck() *MythDeck {
	deck := MythDeck{
		handy.NewDeck([]handy.CardInterface{}),
	}

	return &deck
}
