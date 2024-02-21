package tundra

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// General behavior for a source of commands for the
// command processor to look at.
// Generally, commands should pe placed on the object
// that is taking the action involved, not the object
// the command is acting on. There are exceptions, such
// as an unique command for an object that is not normally
// available, e. g., a fly command for a spaceship,
// or a shoot command for a laser.
type HasCommands interface {
	AddCommand(string, Command)
	RemoveCommand(string)
}

type HasObjects interface {
	AddObject(string, *Object)
	RemoveObject(string)
}
