package party

import (
	"fmt"

	"antique/board"
	"antique/elder"
	"antique/party/deck"
	"antique/player"
)

const (
	ReferenceGateTag = "gates"
	ReferenceClueTag = "clues"
)

var (
	// References Условия усложнения от количества игроков
	References = map[int]map[string]int{
		1: {
			ReferenceGateTag: 1,
			ReferenceClueTag: 1,
		},
	}
)

// Party Партия (игра, катка)
type Party struct {
	board         *board.Board
	elder         *elder.Elder
	investigators []*player.Investigator
	spellDeck     *deck.SpellDeck
	conditionDeck *deck.ConditionDeck
	mythDeck      *deck.MythDeck
	players       []*player.Player
}

func New(numberOfPlayers int, doomTrack int) *Party {
	party := Party{
		board:         board.New(doomTrack, References[numberOfPlayers][ReferenceGateTag], References[numberOfPlayers][ReferenceClueTag]),
		elder:         elder.New(),
		spellDeck:     deck.NewSpellDeck(),
		conditionDeck: deck.NewConditionDeck(),
		mythDeck:      deck.NewMythDeck(),
	}

	firstIsLead := true
	party.investigators = make([]*player.Investigator, numberOfPlayers)
	party.players = make([]*player.Player, numberOfPlayers)
	for i := 0; i < numberOfPlayers; i++ {
		investigator := player.NewInvestigator(firstIsLead)

		party.investigators[i] = investigator
		party.players[i] = player.NewPlayer(investigator)

		firstIsLead = false
	}

	return &party
}

func (p *Party) Process() {
	fmt.Println("начали процесс игры")

	for {
		//	todo фаза действий

		//	todo фаза контакты

		//	todo фаза мифов

		//	todo проверка окончания игры
		if p.checkFinal() {
			fmt.Println("наступил финал партии")

			break
		}
	}

	fmt.Println("закончили процесс игры")
}

func (p *Party) checkFinal() bool {
	// todo переделать на условие от Древнего
	return true
}

func (p *Party) actionPhase() {
	// for _, item := range p.players {
	// 	// item.
	// }
}
