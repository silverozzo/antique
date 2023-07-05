package board

import (
	"antique/board/deck"
	"antique/board/location"
	"antique/board/omen"
)

// Доска, карта местности (стол)
type Board struct {
	gateStack      *deck.GateStack
	clueStack      *deck.ClueStack
	monsterStack   *deck.MonsterStack
	expeditionDeck *deck.ExpeditionDeck
	assetDeck      *deck.AssetDeck
	doomTrack      int
	omen           *omen.Omen
	reserve        [4]*deck.AssetCard
	locations      *location.Graph
}

func New(doomTrack int, startGates int, startClues int) *Board {
	board := Board{
		gateStack:      deck.NewGateStack(),
		clueStack:      deck.NewClueStack(),
		monsterStack:   deck.NewMonsterStack(),
		expeditionDeck: deck.NewExpeditionDeck(),
		assetDeck:      deck.NewAssetDeck(),
		doomTrack:      doomTrack,
		omen:           omen.New(),
		locations:      location.NewGraph(),
	}

	board.OpenGates(startGates)
	board.FillReserves()
	board.UpdateExpedition()
	board.PlaceClues(startClues)

	return &board
}

func (board *Board) OpenGates(count int) {
	for i := 0; i < count; i++ {
		gate := board.gateStack.DiscardTop().(*deck.GateToken)
		monster := board.monsterStack.DiscardRandom().(*deck.MonsterToken)

		location := board.locations.GetLocationByName(gate.GetLocationName())
		location.SetGate(gate)
		location.AddMonster(monster)
	}
}

func (board *Board) FillReserves() {
	for i, item := range board.reserve {
		if item == nil {
			board.reserve[i] = board.assetDeck.DiscardTop().(*deck.AssetCard)
		}
	}
}

func (board *Board) UpdateExpedition() {
	for _, item := range board.locations.GetList() {
		item.UnsetExpedition()
	}

	side := board.expeditionDeck.GetBackSideFromTop()
	locName, ok := side[deck.LocationTag]
	if !ok {
		panic("на рубашке карты экспедиции не определена локация")
	}

	loc := board.locations.GetLocationByName(locName)
	loc.SetExpedition()
}

func (board *Board) PlaceClues(count int) {
	for i := 0; i < count; i++ {
		clue := board.clueStack.DiscardRandom().(*deck.ClueToken)
		loc := board.locations.GetLocationByName(clue.GetLocationName())
		loc.SetClue(clue)
	}
}

func (board *Board) GetNeighborsOfLocation(locName string) []*location.Location {
	loc := board.locations.GetLocationByName(locName)

	return loc.GetNeighbors()
}
