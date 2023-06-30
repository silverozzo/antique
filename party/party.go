package party

import (
	"fmt"

	"antique/board"
	"antique/elder"
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
}

func New() *Party {
	party := Party{
		board: board.New(),
		elder: elder.New(),
	}

	firstIsLead := true
	party.investigators = make([]player.Investigator, NumberOfPlayers)
	for i := 0; i < NumberOfPlayers; i++ {
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
