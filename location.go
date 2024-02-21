package tundra

type Location struct {
	Title       string
	Description string
	Objects     map[string]*Object
	Commands    map[string]Command
}

func (l *Location) AddCommand(name string, command Command) {
	if l.Commands == nil {
		l.Commands = make(map[string]Command, 0)
	}
	l.Commands[name] = command
}

func (l *Location) GetCommand(name string) Command {
	return l.Commands[name]
}

func (l *Location) RemoveCommand(name string) {
	delete(l.Commands, name)
}

func (l *Location) AddObject(name string, object *Object) {
	if l.Objects == nil {
		l.Objects = make(map[string]*Object, 0)
	}
	l.Objects[name] = object
}

func (l *Location) GetObject(name string) *Object {
	return l.Objects[name]
}

func (l *Location) RemoveObject(name string) {
	delete(l.Objects, name)
}
