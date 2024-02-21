package tundra

type CommandResultType int

const (
	Ok   CommandResultType = iota
	Expo                   // display a sequence of messages with a "press enter..." in between
	Death
)

// The first parameter is a list of targets for the command (eg, examine CLOSET, put FISH on TABLE).
// The second is for taking inventory, breaking the world, changing player pos, etc.
type Command func([]*Object, *World) (CommandResults, error)

type CommandResults struct {
	Result CommandResultType
	Msg    []string // unless Result is Expo, only Msg[0] is displayed
}

type CommandProcessor interface {
	UpdateContext()
	Execute(string) (CommandResults, error)
}
