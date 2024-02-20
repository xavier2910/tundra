package tundra

type Player struct {
	CurLoc            *Location
	Inventory         map[string]*Object
	AdditionalContext map[string]*Object
	Commands          map[string]Command
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
	p.Commands[name] = command
}

func (p *Player) RemoveCommand(name string) {
	delete(p.Commands, name)
}

func (p *Player) AddObject(name string, object *Object) {
	p.Inventory[name] = object
}

func (p *Player) RemoveObject(name string) {
	delete(p.Inventory, name)
}
