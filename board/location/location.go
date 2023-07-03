package location

import "antique/handy"

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
	gate      *handy.Card
	monsters  []*handy.Card
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

func (loc *Location) SetGate(gate *handy.Card) {
	loc.gate = gate
}

func (loc *Location) AddMonster(monster *handy.Card) {
	loc.monsters = append(loc.monsters, monster)
}
