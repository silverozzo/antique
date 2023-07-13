package handy

// Card Игральная карта
type Card struct {
	CardInterface

	front Side
	back  Side
}

func NewCard(front Side, back Side) CardInterface {
	return &Card{
		front: front,
		back:  back,
	}
}

func (card *Card) GetFront() Side {
	return card.front
}

func (card *Card) GetBack() Side {
	return card.back
}
