package commands

import (
	"fmt"

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

func Take(name string, obj *tundra.Object, gameData *tundra.World) tundra.Command {
	return func(o []*tundra.Object) (tundra.CommandResults, error) {
		if gameData.PlayerData.Inventory[name] != nil {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg: []string{
					fmt.Sprintf("You already have the %s.", name),
				},
			}, nil
		}
		if gameData.PlayerData.CurLoc.GetObject(name) == nil {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg: []string{
					fmt.Sprintf("You already have the %s, or the %s has ceased to exist.", name, name),
				},
			}, fmt.Errorf("object \"%s\": %#v is nil at location %#v", name, obj, gameData.PlayerData.CurLoc)
		}
		gameData.PlayerData.CurLoc.RemoveObject(name)
		gameData.PlayerData.AddObject(name, obj)
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{fmt.Sprintf("%s taken.", name)},
		}, nil
	}
}

func Drop(name string, obj *tundra.Object, gameData *tundra.World) tundra.Command {
	return func(o []*tundra.Object) (tundra.CommandResults, error) {
		if gameData.PlayerData.Inventory[name] == nil {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg: []string{
					fmt.Sprintf("You don't have a %s.", name),
				},
			}, nil
		}
		if gameData.PlayerData.CurLoc.GetObject(name) != nil {
			return tundra.CommandResults{
				Result: tundra.Ok,
				Msg: []string{
					fmt.Sprintf("There is already a %s here.", name),
				},
			}, fmt.Errorf("duplicate object \"%s\" at location %#v", name, gameData.PlayerData.CurLoc)
		}
		gameData.PlayerData.CurLoc.AddObject(name, obj)
		gameData.PlayerData.RemoveObject(name)
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg:    []string{fmt.Sprintf("%s dropped.", name)},
		}, nil
	}
}
