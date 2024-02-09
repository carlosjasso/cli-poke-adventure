package prompt

type CommandName string

type Command struct {
	Name        CommandName
	Description string
	Action      func() error
}
