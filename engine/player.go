package engine

type Player struct {
	CurLoc            *Location
	Inventory         []*Object
	AdditionalContext []*Object
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

func WithStartingInventory(inventory []*Object) PlayerOption {
	return func(p *Player) {
		p.Inventory = inventory
	}
}

func WithAdditionalContext(context []*Object) PlayerOption {
	return func(p *Player) {
		p.AdditionalContext = context
	}
}

func WithInitialCommands(commands map[string]Command) PlayerOption {
	return func(p *Player) {
		p.Commands = commands
	}
}

func (p *Player) AddCommand(name string, command Command) {
	p.Commands[name] = command
}

func (p *Player) RemoveCommand(name string) {
	delete(p.Commands, name)
}
