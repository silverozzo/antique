package location

import (
	"antique/board/deck"
)

type locationSpace int

const (
	Undefined locationSpace = iota
	City
	Ocean
	Wild
)

type Location struct {
	name      string
	space     locationSpace
	neighbors []*Location
	gate      *deck.GateToken
	monsters  []*deck.MonsterToken
}

func NewLocation(name string, space locationSpace) *Location {
	return &Location{
		name:  name,
		space: space,
	}
}

func (loc *Location) SetNeigbors(neighbors []*Location) {
	loc.neighbors = neighbors
}

func (loc *Location) SetGate(gate *deck.GateToken) {
	loc.gate = gate
}

func (loc *Location) AddMonster(monster *deck.MonsterToken) {
	loc.monsters = append(loc.monsters, monster)
}
