package commandprocessors

import (
	"testing"

	"github.com/xavier2910/tundra"
)

func TestPreprocess(t *testing.T) {

	cp := NewTurnBased(tundra.NewWorld(tundra.NewPlayer(), []*tundra.Location{}))

	t.Run("NulladCommand", func(t *testing.T) {
		owner, cmd, args := cp.preprocess("inventory")
		cmdwant := "inventory"
		if cmd != cmdwant {
			t.Errorf("want command %s, got %s", cmdwant, cmd)
		}
		if len(owner) != 0 {
			t.Errorf("want empty owner, got %s", owner)
		}
		if len(args) != 0 {
			t.Errorf("want no args, got %#v", args)
		}
	})

	t.Run("MonadCommand", func(t *testing.T) {
		owner, cmd, args := cp.preprocess("push button")
		cmdwant := "push"
		ownerwant := "button"
		if cmd != cmdwant {
			t.Errorf("want command %s, got %s", cmdwant, cmd)
		}
		if owner != ownerwant {
			t.Errorf("want owner %s, got %s", ownerwant, owner)
		}
		if len(args) != 0 {
			t.Errorf("want no args, got %#v", args)
		}
	})

	t.Run("PolyadCommand", func(t *testing.T) {
		owner, cmd, args := cp.preprocess("puton table plate cup food")
		cmdwant := "puton"
		ownerwant := "table"
		lenargswant := 3
		if cmd != cmdwant {
			t.Errorf("want command %s, got %s", cmdwant, cmd)
		}
		if owner != ownerwant {
			t.Errorf("want owner %s, got %s", ownerwant, owner)
		}
		if len(args) != lenargswant {
			t.Errorf("want %d args (plate, cup, food), got %d (%#v)", lenargswant, len(args), args)
		}
	})
}

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
	teleporter.AddCommand("examine", func(o []*tundra.Object) (tundra.CommandResults, error) {
		return tundra.CommandResults{
			Msg: []string{
				o[0].Description,
			},
		}, nil
	})
	button.AddCommand("examine", func(o []*tundra.Object) (tundra.CommandResults, error) {
		return tundra.CommandResults{
			Msg: []string{
				o[0].Description,
			},
		}, nil
	})
	fish.AddCommand("flop", func(o []*tundra.Object) (tundra.CommandResults, error) {
		return tundra.CommandResults{Msg: []string{"flop!"}}, nil
	})
	hand.AddCommand("grab", func(o []*tundra.Object) (tundra.CommandResults, error) {
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
					"locationcommandexample": func(o []*tundra.Object) (tundra.CommandResults, error) {
						return tundra.CommandResults{}, nil
					},
				},
			},
		},
	)
	world.PlayerData.CurLoc = world.Places[0]
	world.PlayerData.AddObject("fish", fish)
	world.PlayerData.AddCommand("playercommandexample", func(o []*tundra.Object) (tundra.CommandResults, error) {
		return tundra.CommandResults{}, nil
	})

	cp := NewTurnBased(world)
	cp.UpdateContext()

	t.Run("PlayerInventoryObject", func(t *testing.T) {
		if cp.context["fish"] != nil {
			t.Errorf("loaded player inventory object \"%s\", should not load these!", "fish")
		}
	})
	t.Run("PlayerAdditionalObject", func(t *testing.T) {
		if cp.context["hand"] == nil {
			t.Errorf("failed to load player additional context object command")
		}
	})
	t.Run("LocationObject", func(t *testing.T) {
		if cp.context["teleporter"] == nil {
			t.Errorf("failed to load location object command")
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
	teleporter.AddCommand("examine", func(o []*tundra.Object) (tundra.CommandResults, error) {
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg: []string{
				teleporter.Description,
			},
		}, nil
	})

	button.AddCommand("examine", func(o []*tundra.Object) (tundra.CommandResults, error) {
		return tundra.CommandResults{
			Result: tundra.Ok,
			Msg: []string{
				button.Description,
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

		cp.clearContext()

		cp.addContext(world.PlayerData.CurLoc.Objects)

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
