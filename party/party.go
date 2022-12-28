package party

import "fmt"

type Party struct{}

func New() *Party {
	return &Party{}
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
