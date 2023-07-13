package handy

// Deck Колода игральных карт
type Deck struct {
	DeckInterface

	cards []CardInterface
}

func NewDeck(cards []CardInterface) DeckInterface {
	return &Deck{
		cards: cards,
	}
}

func (deck *Deck) SetCards(cards []CardInterface) {
	deck.cards = cards
}

func (deck *Deck) AddCardToBottom(card CardInterface) {
	deck.cards = append(deck.cards, card)
}

func (deck *Deck) DiscardTop() CardInterface {
	if len(deck.cards) == 0 {
		panic("Колода пустая")
	}

	discarded := deck.cards[0]
	deck.cards = deck.cards[1:]

	return discarded
}

func (deck *Deck) GetBackSideFromTop() Side {
	if len(deck.cards) == 0 {
		panic("Колода пустая")
	}

	return deck.cards[0].GetBack()
}
