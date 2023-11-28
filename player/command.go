package player

type Command struct {
	action    ActionType
	arguments map[string]interface{}
}

func NewCommand(action ActionType, arguments map[string]interface{}) *Command {
	comm := Command{
		action:    action,
		arguments: arguments,
	}

	return &comm
}
