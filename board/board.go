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
	doomTrack      int
	omen           *omen.Omen
	reserve        [4]*deck.AssetCard
	locations      *location.Graph
}

func New(doomTrack int) *Board {
	board := Board{
		gateStack:      deck.NewGateStack(),
		clueStack:      deck.NewClueStack(),
		monsterStack:   deck.NewMonsterStack(),
		expeditionDeck: deck.NewExpeditionDeck(),
		doomTrack:      doomTrack,
		omen:           &omen.Omen{},
		locations:      location.NewGraph(),
	}

	// делаем пока примитивно-упрощенно
	board.OpenTopGate()

	return &board
}

// Открываем ворота с верху стопки врат
func (board *Board) OpenTopGate() {
	// Берем верхнюю карту со стопки врат
	// gate := board.gateStack.DiscardTop()
	// monster := board.monsterStack.DiscardTop()

	// location := board.locations.GetLocationByName(gate.GetFront().GetLocationName())

	// panic("not implemented yet")
}
