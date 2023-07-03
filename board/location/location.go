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
	name          string
	space         locationSpace
	neighbors     []*Location
	gate          *deck.GateToken
	monsters      []*deck.MonsterToken
	hasExpedition bool
	clue          *deck.ClueToken
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

func (loc *Location) SetExpedition() {
	loc.hasExpedition = true
}

func (loc *Location) UnsetExpedition() {
	loc.hasExpedition = false
}

func (loc *Location) HasExpedition() bool {
	return loc.hasExpedition
}

func (loc *Location) SetClue(clue *deck.ClueToken) {
	loc.clue = clue
}

func (loc *Location) DiscardClue() *deck.ClueToken {
	clue := loc.clue
	loc.clue = nil

	return clue
}
