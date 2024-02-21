package commands

import (
	"github.com/xavier2910/tundra"
)

func Examine(owner *tundra.Object) tundra.Command {
	return func(o []*tundra.Object) (tundra.CommandResults, error) {
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg: []string{
				owner.Description,
			},
		}, nil
	}
}
