package handy

type Stack struct {
	StackInterface

	tokens []TokenInterface
}

func NewStack(tokens []TokenInterface) StackInterface {
	return &Stack{
		tokens: tokens,
	}
}

func (stack *Stack) DiscardTop() TokenInterface {
	if len(stack.tokens) == 0 {
		panic("стопка жетонов пустая")
	}

	discarded := stack.tokens[0]
	stack.tokens = stack.tokens[1:]

	return discarded
}

func (stack *Stack) DiscardRandom() TokenInterface {
	if len(stack.tokens) == 0 {
		panic("стопка жетонов пустая")
	}

	//	пока из случайного возвращаем верхний
	discarded := stack.tokens[0]
	stack.tokens = stack.tokens[1:]

	return discarded
}
