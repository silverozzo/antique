package deck

import (
	"antique/handy"
)

type GateToken struct {
	handy.TokenInterface
}

type GateStack struct {
	handy.StackInterface
}

func newGate(locName string) *GateToken {
	gate := GateToken{
		handy.NewToken(handy.Side{LocationTag: locName}, handy.Side{}),
	}

	return &gate
}

// Возвращает колоду врат
func NewGateStack() *GateStack {

	// TODO пока хрен знает как сделать... болваночно
	one := newGate("Лондон")

	stack := GateStack{
		handy.NewStack([]handy.TokenInterface{one}),
	}

	return &stack
}
