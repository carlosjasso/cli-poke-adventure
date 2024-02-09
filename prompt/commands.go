package prompt

const (
	COMMAND_HELP CommandName = "help"
	COMMAND_EXIT CommandName = "exit"
)

var Commands = map[CommandName]*Command{
	COMMAND_HELP: {
		Name:        COMMAND_HELP,
		Description: "Some help",
		Action:      printHelp,
	},
	COMMAND_EXIT: {
		Name:        COMMAND_EXIT,
		Description: "",
		Action:      nil,
	},
}
