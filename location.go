package tundra

import "fmt"

type Direction string

const (
	North     Direction = "n"
	Northeast Direction = "ne"
	East      Direction = "e"
	Southeast Direction = "se"
	South     Direction = "s"
	Southwest Direction = "sw"
	West      Direction = "w"
	Northwest Direction = "nw"
	Up        Direction = "u"
	Down      Direction = "d"
	In        Direction = "in"
	Out       Direction = "ot"
)

// A logical place in the virtual universe of the game.
// In the real world, one can go from one location to
// another (typically). Now, this structure doesn't have
// and explicit connection to other locations. These
// are provided by special objects, called "north",
// "south", etc.
type Location struct {
	Title       string
	Description string
	Objects     map[string]*Object
	Commands    map[string]Command
}

func (l *Location) Describe() (description string) {
	description = l.Description
	for oname := range l.Objects {
		description = fmt.Sprintf("%s. There is a %s here.", description, oname)
	}
	return
}

// Connect the location to another via a named object.
func (l *Location) SetConnection(direction Direction, other *Location, forPlayer *Player, withCP CommandProcessor) {

	if l.Objects == nil {
		l.Objects = make(map[string]*Object, 1)
	}

	// notice that a direction is just a particular object
	// with a go command. "go east" gets resolved to the "go"
	// command on the "east" object by the command processor.

	conn := NewObject()
	conn.AddCommand("go", func(o []*Object) (CommandResults, error) {
		forPlayer.CurLoc = other
		withCP.UpdateContext()
		return CommandResults{
			Result: Ok,
			Msg:    []string{"# " + other.Title + "\n\n" + other.Describe()},
		}, nil
	})
	l.Objects[string(direction)] = conn
}

func (l *Location) RemoveConnection(direction Direction) {
	l.Objects[string(direction)] = nil
}

func (l *Location) AddCommand(name string, command Command) {
	if l.Commands == nil {
		l.Commands = make(map[string]Command, 0)
	}
	l.Commands[name] = command
}

func (l *Location) GetCommand(name string) Command {
	if l.Commands == nil {
		return nil
	}
	return l.Commands[name]
}

func (l *Location) RemoveCommand(name string) {
	if l.Commands == nil {
		return
	}
	delete(l.Commands, name)
}

func (l *Location) AddObject(name string, object *Object) {
	if l.Objects == nil {
		l.Objects = make(map[string]*Object, 0)
	}
	l.Objects[name] = object
}

func (l *Location) GetObject(name string) *Object {
	if l.Objects == nil {
		return nil
	}
	return l.Objects[name]
}

func (l *Location) RemoveObject(name string) {
	if l.Objects == nil {
		return
	}
	delete(l.Objects, name)
}
