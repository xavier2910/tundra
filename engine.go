package tundra

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type HasCommands interface {
	AddCommand(string, Command)
	RemoveCommand(string)
}

type HasObjects interface {
	AddObject(string, *Object)
	RemoveObject(string)
}
