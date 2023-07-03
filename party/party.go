package party

import (
	"fmt"

	"antique/board"
	"antique/elder"
	"antique/party/deck"
	"antique/player"
)

const (
	NumberOfPlayers = 1
)

// Партия (игра, катка)
type Party struct {
	board         *board.Board
	elder         *elder.Elder
	investigators []player.Investigator
	spellDeck     *deck.SpellDeck
	conditionDeck *deck.ConditionDeck
	mythDeck      *deck.MythDeck
}

func New(numberOfPlayers int, doomTrack int) *Party {
	party := Party{
		board:         board.New(doomTrack),
		elder:         elder.New(),
		spellDeck:     deck.NewSpellDeck(),
		conditionDeck: deck.NewConsitionDeck(),
		mythDeck:      deck.NewMythDeck(),
	}

	firstIsLead := true
	party.investigators = make([]player.Investigator, numberOfPlayers)
	for i := 0; i < numberOfPlayers; i++ {
		party.investigators[i] = *player.NewInvestigator(firstIsLead)

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
	return true
}
