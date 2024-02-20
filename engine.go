package tundra

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type CommandResultType int

const (
	Ok   CommandResultType = iota
	Expo                   // display a sequence of messages with a "press enter..." in between
	Death
)

type Command func(*World) (CommandResults, error)

type CommandResults struct {
	Result CommandResultType
	Msg    []string // unless Result is Expo, only Msg[0] is displayed
}

type HasCommands interface {
	AddCommand(string, Command)
	RemoveCommand(string)
}

type HasObjects interface {
	AddObject(string, *Object)
	RemoveObject(string)
}
