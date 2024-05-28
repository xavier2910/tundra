package commandprocessors

import (
	"fmt"
	"strings"

	"github.com/xavier2910/tundra"
)

type turnBased struct {
	context  map[string]*tundra.Object
	universe *tundra.World
}

// A CommandProcessor impementation which updates the gamestate in a
// turn-based format (updates world once per each time Execute() is
// called).
func NewTurnBased(wrld *tundra.World) *turnBased {
	return &turnBased{universe: wrld}
}

// Updates object and command context in priority order:
// player objects, location objects, (rest commands only) player, location.
func (cp *turnBased) UpdateContext() {
	cp.clearContext()

	cp.appendContext(cp.universe.PlayerData.AdditionalContext)
	cp.appendContext(cp.universe.PlayerData.CurLoc.Objects)
}

// overwrites in case of duplicate
func (cp *turnBased) InjectContext(name string, item *tundra.Object) {
	if cp.context == nil {
		cp.context = make(map[string]*tundra.Object, 1)
	}
	cp.context[name] = item
}

func (cp *turnBased) clearContext() {
	cp.context = make(map[string]*tundra.Object, 0)
}

// Adds the given objects to the search path of the
// CommandProcessor, overwriting any existing
// commands if a naming conflict arises.
func (cp *turnBased) addContext(additionalContext map[string]*tundra.Object) {

	for key, obj := range additionalContext {
		cp.context[key] = obj
	}
}

// Adds the given objects to the search path of the
// CommandProcessor, skipping any new commands
// if a naming conflict arises.
func (cp *turnBased) appendContext(additionalContext map[string]*tundra.Object) {

	for key, obj := range additionalContext {

		if cp.context[key] == nil {
			cp.context[key] = obj
		}
	}
}

func (cp *turnBased) Execute(command string) (tundra.CommandResults, error) {

	owner, cmd, args := cp.preprocess(command)
	parsedcmd, targets := cp.resolve(owner, cmd, args)

	if parsedcmd != nil {
		return parsedcmd(targets)
	} else {
		return tundra.CommandResults{Msg: []string{fmt.Sprintf("You cannot %s the %s.", cmd, owner)}}, fmt.Errorf("cmd '%s' is not available", command)
	}
}

func (cp *turnBased) preprocess(input string) (owner, cmd string, targets []string) {

	split := strings.Fields(input)
	switch len(split) {
	case 0:
		owner = ""
		cmd = ""
		targets = []string{}
	case 1:
		owner = ""
		cmd = split[0]
		targets = []string{}
	case 2:
		owner = split[1]
		cmd = split[0]
		targets = []string{}
	default:
		owner = split[1]
		cmd = split[0]
		targets = split[2:]
	}
	return
}

func (cp *turnBased) resolve(owner, cmd string, args []string) (command tundra.Command, arguments []*tundra.Object) {

	if len(owner) != 0 {
		parsedOwner := cp.context[owner]
		if parsedOwner != nil {
			command = parsedOwner.GetCommand(cmd)
		} else {
			command = func(o []*tundra.Object) (tundra.CommandResults, error) {
				return tundra.CommandResults{
						Result: tundra.Ok,
						Msg:    []string{fmt.Sprintf("There is no %s here.", owner)},
					},
					fmt.Errorf("could not resolve owner object %s", owner)
			}
		}

	} else if maybe := cp.universe.PlayerData.GetCommand(cmd); maybe != nil {
		command = maybe

	} else if maybe := cp.universe.PlayerData.CurLoc.GetCommand(cmd); maybe != nil {
		command = maybe

	} else {
		command = func(o []*tundra.Object) (tundra.CommandResults, error) {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg:    []string{fmt.Sprintf("Can't apply %s in this context.", cmd)},
			}, fmt.Errorf("command %s not found", cmd)
		}
	}

	arguments = make([]*tundra.Object, len(args))
	for i, arg := range args {
		arguments[i] = cp.context[arg]
	}

	return
}
