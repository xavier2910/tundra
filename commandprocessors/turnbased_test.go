package commandprocessors

import (
	"testing"

	"github.com/xavier2910/tundra"
)

func TestUpdateContext(t *testing.T) {

	teleporterDesc := "The teleporter is equipped with a single button."

	teleporter := tundra.NewObject(
		tundra.WithDescription(teleporterDesc),
	)
	button := tundra.NewObject(
		tundra.WithDescription("Shiny button"),
	)
	fish := tundra.NewObject()
	hand := tundra.NewObject()

	teleporter.AddObject("button", button)
	teleporter.AddCommand("examine", func(o []*tundra.Object, w *tundra.World) (tundra.CommandResults, error) {
		return tundra.CommandResults{
			Msg: []string{
				o[0].Description,
			},
		}, nil
	})
	button.AddCommand("examine", func(o []*tundra.Object, w *tundra.World) (tundra.CommandResults, error) {
		return tundra.CommandResults{
			Msg: []string{
				o[0].Description,
			},
		}, nil
	})
	fish.AddCommand("flop", func(o []*tundra.Object, w *tundra.World) (tundra.CommandResults, error) {
		return tundra.CommandResults{Msg: []string{"flop!"}}, nil
	})
	hand.AddCommand("grab", func(o []*tundra.Object, w *tundra.World) (tundra.CommandResults, error) {
		return tundra.CommandResults{}, nil
	})

	world := tundra.NewWorld(
		tundra.NewPlayer(
			tundra.WithAdditionalContext(map[string]*tundra.Object{
				"hand": hand,
			}),
		),
		[]*tundra.Location{
			{
				Title:       "Cave Entrance",
				Description: "You are looking into the mouth of a dark cave. There is also a teleporter nearby.",
				Objects: map[string]*tundra.Object{
					"teleporter": teleporter,
				},
				Commands: map[string]tundra.Command{
					"locationcommandexample": func(o []*tundra.Object, w *tundra.World) (tundra.CommandResults, error) {
						return tundra.CommandResults{}, nil
					},
				},
			},
		},
	)
	world.PlayerData.CurLoc = world.Places[0]
	world.PlayerData.AddObject("fish", fish)
	world.PlayerData.AddCommand("playercommandexample", func(o []*tundra.Object, w *tundra.World) (tundra.CommandResults, error) {
		return tundra.CommandResults{}, nil
	})

	cp := NewTurnBased(world)
	cp.UpdateContext()

	t.Run("PlayerInventoryObject", func(t *testing.T) {
		if cp.commandContext["flop"] == nil {
			t.Errorf("failed to load player inventory object command")
		}
	})
	t.Run("PlayerAdditionalObject", func(t *testing.T) {
		if cp.commandContext["grab"] == nil {
			t.Errorf("failed to load player additional context object command")
		}
	})
	t.Run("LocationObject", func(t *testing.T) {
		if cp.commandContext["examine"] == nil {
			t.Errorf("failed to load location object command")
		}
	})
	t.Run("Player", func(t *testing.T) {
		if cp.commandContext["playercommandexample"] == nil {
			t.Errorf("failed to load player command")
		}
	})
	t.Run("Location", func(t *testing.T) {
		if cp.commandContext["locationcommandexample"] == nil {
			t.Errorf("failed to load location command")
		}
	})
}

func TestExecute(t *testing.T) {
	world := tundra.NewWorld(
		tundra.NewPlayer(),
		[]*tundra.Location{},
	)

	cp := NewTurnBased(world)

	t.Run("Nothing", func(t *testing.T) {
		result, err := cp.Execute("go east")

		if err == nil {
			t.Errorf("Bad command got no error (result = %#v)", result)
		}
	})

	t.Run("Examine", testExecuteExamine(t, world, cp))
}

func testExecuteExamine(t *testing.T, world *tundra.World, cp *turnBased) func(*testing.T) {
	t.Helper()

	teleporterDesc := "The teleporter is equipped with a single button."

	teleporter := tundra.NewObject(
		tundra.WithDescription(teleporterDesc),
	)
	button := tundra.NewObject(
		tundra.WithDescription("Shiny button"),
	)

	teleporter.AddObject("button", button)
	teleporter.AddCommand("examine", func(o []*tundra.Object, w *tundra.World) (tundra.CommandResults, error) {
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg: []string{
				o[0].Description,
			},
		}, nil
	})

	button.AddCommand("examine", func(o []*tundra.Object, w *tundra.World) (tundra.CommandResults, error) {
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg: []string{
				o[0].Description,
			},
		}, nil
	})

	world.Places = []*tundra.Location{
		{
			Title:       "Cave Entrance",
			Description: "You are looking into the mouth of a dark cave. There is also a teleporter nearby.",
			Objects: map[string]*tundra.Object{
				"teleporter": teleporter,
			},
		},
	}
	world.PlayerData = tundra.NewPlayer(tundra.WithStartingLocation(world.Places[0]))

	return func(t *testing.T) {

		cp.clearCommandContext()
		cp.clearObjectContext()

		for _, obj := range world.PlayerData.CurLoc.Objects {
			cp.addCommandContext(obj.Commands)
		}
		cp.addObjectContext(world.PlayerData.CurLoc.Objects)

		result, err := cp.Execute("examine teleporter")

		if err != nil || result.Result != tundra.Ok || result.Msg[0] != teleporterDesc {
			t.Errorf("want error nil and result %#v, got error %#v and result %#v",
				tundra.CommandResults{
					Result: tundra.Ok,
					Msg:    []string{teleporterDesc},
				},
				err,
				result,
			)
		}
	}
}
