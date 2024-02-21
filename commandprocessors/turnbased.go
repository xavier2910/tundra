package commandprocessors

import (
	"fmt"
	"strings"

	"github.com/xavier2910/tundra"
)

// A CommandProcessor impementation which updates the gamestate in a
// turn-based format (updates world once per each time Execute() is
// called)
type turnBased struct {
	commandContext map[string]tundra.Command
	objectContext  map[string]*tundra.Object
	universe       *tundra.World
}

func NewTurnBased(wrld *tundra.World) *turnBased {
	return &turnBased{universe: wrld}
}

// Updates object and command context in priority order:
// player objects, location objects, (rest commands only) player, location.
func (cp *turnBased) UpdateContext() {
	cp.clearCommandContext()
	cp.clearObjectContext()

	cp.appendObjectContext(cp.universe.PlayerData.Inventory)
	cp.appendObjectContext(cp.universe.PlayerData.AdditionalContext)
	cp.appendObjectContext(cp.universe.PlayerData.CurLoc.Objects)

	cp.appendCommandContext(cp.universe.PlayerData.Commands)
	cp.appendCommandContext(cp.universe.PlayerData.CurLoc.Commands)
}

func (cp *turnBased) clearCommandContext() {
	cp.commandContext = make(map[string]tundra.Command, 0)
}

// Adds the given commands to the search path of the
// CommandProcessor, overwriting any existing
// commands if a naming conflict arises.
func (cp *turnBased) addCommandContext(additionalContext map[string]tundra.Command) {

	for key, cmd := range additionalContext {
		cp.commandContext[key] = cmd
	}
}

// Adds the given commands to the search path of the
// CommandProcessor, skipping any new commands
// if a naming conflict arises.
func (cp *turnBased) appendCommandContext(additionalContext map[string]tundra.Command) {

	for key, cmd := range additionalContext {

		if cp.commandContext[key] == nil {
			cp.commandContext[key] = cmd
		}
	}
}

func (cp *turnBased) clearObjectContext() {
	cp.objectContext = make(map[string]*tundra.Object, 0)
}

// Adds the given objects to the search path of the
// CommandProcessor, overwriting any existing
// commands if a naming conflict arises.
func (cp *turnBased) addObjectContext(additionalContext map[string]*tundra.Object) {

	for key, obj := range additionalContext {
		cp.objectContext[key] = obj
		cp.addCommandContext(obj.Commands)
	}
}

// Adds the given objects to the search path of the
// CommandProcessor, skipping any new commands
// if a naming conflict arises.
func (cp *turnBased) appendObjectContext(additionalContext map[string]*tundra.Object) {

	for key, obj := range additionalContext {

		if cp.objectContext[key] == nil {
			cp.objectContext[key] = obj
			cp.appendCommandContext(obj.Commands)
		}
	}
}

func (cp *turnBased) Execute(command string) (tundra.CommandResults, error) {

	cmd, args := cp.preprocess(command)
	parsedcmd, targets := cp.resolve(cmd, args)

	if parsedcmd != nil {
		return parsedcmd(targets, cp.universe)
	} else {
		return tundra.CommandResults{Msg: []string{"bad cmd"}}, fmt.Errorf("cmd '%s' is not available", command)
	}
}

func (cp *turnBased) preprocess(cmd string) (string, []string) {

	split := strings.Fields(cmd)

	return split[0], split[1:]
}

func (cp *turnBased) resolve(cmd string, args []string) (tundra.Command, []*tundra.Object) {

	arguments := make([]*tundra.Object, 2)

	for i, arg := range args {
		arguments[i] = cp.objectContext[arg]
	}

	return cp.commandContext[cmd], arguments
}
