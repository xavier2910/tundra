package tundra

type Player struct {
	CurLoc            *Location
	Inventory         map[string]*Object
	AdditionalContext map[string]*Object
	commands          map[string]Command
}
type PlayerOption func(*Player)

func NewPlayer(options ...PlayerOption) *Player {
	plyr := &Player{}

	for _, option := range options {
		option(plyr)
	}

	return plyr
}

func WithStartingLocation(l *Location) PlayerOption {
	return func(p *Player) {
		p.CurLoc = l
	}
}

func WithStartingInventory(inventory map[string]*Object) PlayerOption {
	return func(p *Player) {
		p.Inventory = inventory
	}
}

func WithAdditionalContext(context map[string]*Object) PlayerOption {
	return func(p *Player) {
		p.AdditionalContext = context
	}
}

func (p *Player) AddCommand(name string, command Command) {
	if p.commands == nil {
		p.commands = make(map[string]Command, 0)
	}
	p.commands[name] = command
}

func (p *Player) GetCommand(name string) Command {
	if p.commands == nil {
		return nil
	}
	return p.commands[name]
}

func (p *Player) RemoveCommand(name string) {
	if p.commands == nil {
		return
	}
	delete(p.commands, name)
}

// adds to inventory
func (p *Player) AddObject(name string, object *Object) {
	if p.Inventory == nil {
		p.Inventory = make(map[string]*Object, 0)
	}
	p.Inventory[name] = object
}

// adds to inventory
func (p *Player) RemoveObject(name string) {
	delete(p.Inventory, name)
}
