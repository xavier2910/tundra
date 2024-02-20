package commands

import (
	"github.com/xavier2910/tundra"
)

func Examine() tundra.Command {
	return func(o []*tundra.Object, w *tundra.World) (tundra.CommandResults, error) {
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg: []string{
				o[0].Description,
			},
		}, nil
	}
}
