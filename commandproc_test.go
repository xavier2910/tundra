package tundra

import "testing"

func TestExecute(t *testing.T) {
	world := NewWorld(
		NewPlayer(),
		[]*Location{},
	)

	cp := NewCommandProcessor(world)

	t.Run("Nothing", func(t *testing.T) {
		result, err := cp.Execute("go east")

		if err == nil {
			t.Errorf("Bad command got no error (result = %#v)", result)
		}
	})

	t.Run("Examine", testExecuteExamine(t, world, cp))
}

func testExecuteExamine(t *testing.T, world *World, cp *CommandProcessor) func(*testing.T) {
	t.Helper()

	teleporterDesc := "The teleporter is equipped with a single button."

	teleporter := NewObject(
		WithDescription(teleporterDesc),
	)
	button := NewObject(
		WithDescription("Shiny button"),
	)

	teleporter.AddObject("button", button)
	teleporter.AddCommand("examine", func(o []*Object, w *World) (CommandResults, error) {
		return CommandResults{
			Result: Ok,
			Msg: []string{
				o[0].Description,
			},
		}, nil
	})

	button.AddCommand("examine", func(o []*Object, w *World) (CommandResults, error) {
		return CommandResults{
			Result: Ok,
			Msg: []string{
				o[0].Description,
			},
		}, nil
	})

	world.Places = []*Location{
		{
			Title:       "Cave Entrance",
			Description: "You are looking into the mouth of a dark cave. There is also a teleporter nearby.",
			Objects: map[string]*Object{
				"teleporter": teleporter,
			},
		},
	}
	world.PlayerData = NewPlayer(WithStartingLocation(world.Places[0]))

	return func(t *testing.T) {

		cp.ClearCommandContext()
		cp.ClearObjectContext()

		for _, obj := range world.PlayerData.CurLoc.Objects {
			cp.AddCommandContext(obj.Commands)
		}
		cp.AddObjectContext(world.PlayerData.CurLoc.Objects)

		result, err := cp.Execute("examine teleporter")

		if err != nil || result.Result != Ok || result.Msg[0] != teleporterDesc {
			t.Errorf("want error nil and result %#v, got error %#v and result %#v",
				CommandResults{
					Result: Ok,
					Msg:    []string{teleporterDesc},
				},
				err,
				result,
			)
		}
	}
}
