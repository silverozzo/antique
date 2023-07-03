package handy

type Token struct {
	TokenInterface

	front Side
	back  Side
}

func NewToken(front Side, back Side) TokenInterface {
	return &Token{
		front: front,
		back:  back,
	}
}

func (token *Token) GetFront() Side {
	return token.front
}

func (token *Token) GetBack() Side {
	return token.back
}
