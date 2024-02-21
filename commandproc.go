package tundra

// The first parameter is a list of targets for the command (eg, put fish on TABLE,
// NOT examine CLOSET. In the latter case a particular examine implementation has
// already been selected.)
// Any other data that needs to be referenced shoud be closured in command creation.
type Command func([]*Object) (CommandResults, error)

type CommandResults struct {
	Result CommandResultType
	Msg    []string // unless Result is Expo, only Msg[0] is displayed
}

type CommandResultType int

const (
	Ok    CommandResultType = iota // display a single message
	Expo                           // display a sequence of messages with a "press enter..." in between
	Death                          // end the game...
	Win                            // also end the game, but with different results.
)

// The command processor handles taking the user
// input, finding the command and referenced objects,
// and executing the command in question. Commands are
// to be owned by the object they concern. For example,
// an examine command is owned by each object, as each
// reports a different description. Taking inventory
// belongs to the player. It is the command processor's
// task to resolve the correct command to use and it's
// arguments.
type CommandProcessor interface {
	// Load visible targets for command parsing.
	// Should load only things immediately visible
	// to the player.
	UpdateContext()
	// Add additional targets for command parsing.
	// Allows for, say, an object only to be usable
	// on closer inspection or button press, etc.
	InjectContext(string, *Object)

	Execute(string) (CommandResults, error)
}
