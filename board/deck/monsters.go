package deck

import "antique/handy"

type MonsterToken struct {
	handy.TokenInterface
}

type MonsterStack struct {
	handy.StackInterface
}

func newMonster(name string) *MonsterToken {
	return &MonsterToken{
		handy.NewToken(handy.Side{MonsterNameTag: name}, handy.Side{}),
	}
}

func NewMonsterStack() *MonsterStack {
	one := newMonster("Культист")
	deck := MonsterStack{
		handy.NewStack([]handy.TokenInterface{one}),
	}

	return &deck
}
