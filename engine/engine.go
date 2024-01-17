package engine

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

type World struct {
	PlayerData *Player
	Places     []*Location
}

func NewWorld(plyr *Player, places []*Location) *World {
	return &World{
		PlayerData: plyr,
		Places:     places,
	}
}

type Object struct {
	Keyword       string
	Description   string
	ContainedObjs []*Object
	Commands      []Command
}

type Location struct {
	Title              string
	Description        string
	Objects            []*Object
	ConnectedLocations map[Direction]*Location
	Commands           []Command
}
