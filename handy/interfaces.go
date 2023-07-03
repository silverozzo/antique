package handy

type CardInterface interface {
	GetFront() Side
	GetBack() Side
}

type TokenInterface interface {
	GetFront() Side
	GetBack() Side
}

type DeckInterface interface {
	SetCards([]CardInterface)
	AddCardToBottom(CardInterface)
	DiscardTop() CardInterface
}

type StackInterface interface {
	DiscardTop() TokenInterface
	DiscardRandom() TokenInterface
}
