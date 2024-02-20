package tundra

type Location struct {
	Title              string
	Description        string
	Objects            map[string]*Object
	ConnectedLocations map[Direction]*Location
	Commands           map[string]Command
}

func (l *Location) AddConnection(dir Direction, loc *Location) {
	l.ConnectedLocations[dir] = loc
}

func (l *Location) RemoveConnection(dir Direction) {
	delete(l.ConnectedLocations, dir)
}

func (l *Location) AddCommand(name string, command Command) {
	l.Commands[name] = command
}

func (l *Location) RemoveCommand(name string) {
	delete(l.Commands, name)
}

func (l *Location) AddObject(name string, object *Object) {
	l.Objects[name] = object
}

func (l *Location) RemoveObject(name string) {
	delete(l.Objects, name)
}
