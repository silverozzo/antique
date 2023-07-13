package deck

import "antique/handy"

// SpellCard Тип карт заклинаний
type SpellCard struct {
	handy.CardInterface
}

// SpellDeck Колода карт закдлинаний в запасе
type SpellDeck struct {
	handy.DeckInterface
}

func NewSpellDeck() *SpellDeck {
	deck := SpellDeck{
		handy.NewDeck([]handy.CardInterface{}),
	}

	return &deck
}
