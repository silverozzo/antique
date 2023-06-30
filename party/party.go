package party

import (
	"fmt"

	"antique/board"
)

// Партия (игра, катка)
type Party struct {
	board *board.Board
}

func New() *Party {
	return &Party{
		board: board.New(),
	}
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
