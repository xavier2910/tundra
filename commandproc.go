package tundra

import (
	"fmt"
	"strings"
)

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

type CommandProcessor struct {
	commandContext map[string]Command
	objectContext  map[string]*Object
	universe       *World
}

func NewCommandProcessor(wrld *World) *CommandProcessor {
	return &CommandProcessor{universe: wrld}
}

func (cp *CommandProcessor) ClearCommandContext() {
	cp.commandContext = make(map[string]Command, 0)
}

// Adds the given commands to the seach path of the
// CommandProcessor, overwriting any existing
// commands if a naming conflict arises.
func (cp *CommandProcessor) AddCommandContext(additionalContext map[string]Command) {

	for key, cmd := range additionalContext {
		cp.commandContext[key] = cmd
	}
}

// Adds the given commands to the seach path of the
// CommandProcessor, skipping any new commands
// if a naming conflict arises.
func (cp *CommandProcessor) AppendCommandContext(additionalContext map[string]Command) {

	for key, cmd := range additionalContext {

		if cp.commandContext[key] == nil {
			cp.commandContext[key] = cmd
		}
	}
}

func (cp *CommandProcessor) ClearObjectContext() {
	cp.objectContext = make(map[string]*Object, 0)
}

// Adds the given commands to the seach path of the
// CommandProcessor, overwriting any existing
// commands if a naming conflict arises.
func (cp *CommandProcessor) AddObjectContext(additionalContext map[string]*Object) {

	for key, obj := range additionalContext {
		cp.objectContext[key] = obj
	}
}

// Adds the given commands to the seach path of the
// CommandProcessor, skipping any new commands
// if a naming conflict arises.
func (cp *CommandProcessor) AppendObjectContext(additionalContext map[string]*Object) {

	for key, obj := range additionalContext {

		if cp.objectContext[key] == nil {
			cp.objectContext[key] = obj
		}
	}
}

func (cp *CommandProcessor) Execute(command string) (CommandResults, error) {

	cmd, args := cp.preprocess(command)
	parsedcmd, targets := cp.resolve(cmd, args)

	if parsedcmd != nil {
		return parsedcmd(targets, cp.universe)
	} else {
		return CommandResults{Msg: []string{"bad cmd"}}, fmt.Errorf("cmd '%s' is not available", command)
	}
}

func (cp *CommandProcessor) preprocess(cmd string) (string, []string) {

	split := strings.Fields(cmd)

	return split[0], split[1:]
}

func (cp *CommandProcessor) resolve(cmd string, args []string) (Command, []*Object) {

	arguments := make([]*Object, 2)

	for i, arg := range args {
		arguments[i] = cp.objectContext[arg]
	}

	return cp.commandContext[cmd], arguments
}
