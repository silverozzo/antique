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

func New(doomTrack int) *Board {
	board := Board{
		gateStack:      deck.NewGateStack(),
		clueStack:      deck.NewClueStack(),
		monsterStack:   deck.NewMonsterStack(),
		expeditionDeck: deck.NewExpeditionDeck(),
		assetDeck:      deck.NewAssetDeck(),
		doomTrack:      doomTrack,
		omen:           &omen.Omen{},
		locations:      location.NewGraph(),
	}

	board.OpenTopGate()
	board.FillReserves()

	return &board
}

// Открываем ворота с верху стопки врат
func (board *Board) OpenTopGate() {
	gate := board.gateStack.DiscardTop().(*deck.GateToken)
	monster := board.monsterStack.DiscardRandom().(*deck.MonsterToken)

	location := board.locations.GetLocationByName(gate.GetLocationName())
	location.SetGate(gate)
	location.AddMonster(monster)
}

func (board *Board) FillReserves() {
	for i, item := range board.reserve {
		if item == nil {
			board.reserve[i] = board.assetDeck.DiscardTop().(*deck.AssetCard)
		}
	}
}
