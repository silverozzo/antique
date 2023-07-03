package deck

import "antique/handy"

type ClueToken struct {
	handy.TokenInterface
}

type ClueStack struct {
	handy.StackInterface
}

func newClue(locName string) *ClueToken {
	clue := ClueToken{
		handy.NewToken(handy.Side{LocationTag: locName}, handy.Side{}),
	}

	return &clue
}

func NewClueStack() *ClueStack {
	one := newClue("Побережье")
	stack := ClueStack{
		handy.NewStack([]handy.TokenInterface{one}),
	}

	return &stack
}

func (clue *ClueToken) GetLocationName() string {
	return clue.GetFront()[LocationTag]
}
