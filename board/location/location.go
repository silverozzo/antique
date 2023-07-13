package location

import (
	"antique/board/deck"
)

type SpaceType int

const (
	City SpaceType = iota
	Sea
	Wild
)

type Location struct {
	name          string
	space         SpaceType
	gate          *deck.GateToken
	monsters      []*deck.MonsterToken
	hasExpedition bool
	clue          *deck.ClueToken
	routes        []*Route
}

func NewLocation(name string, space SpaceType) *Location {
	return &Location{
		name:  name,
		space: space,
	}
}

func (loc *Location) SetRoutes(routes []*Route) {
	for _, i := range routes {
		if i.locations[0] != loc && i.locations[1] != loc {
			panic("в локацию устанавливают пути, которые не связаны с текущей")
		}
	}

	loc.routes = routes
}

func (loc *Location) GetName() string {
	return loc.name
}

func (loc *Location) GetNeighbors() []*Location {
	var neighbors []*Location

	for _, i := range loc.routes {
		neighbors = append(neighbors, i.GetAnotherEnd(loc))
	}

	return neighbors
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
