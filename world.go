package tundra

type World struct {
	PlayerData *Player
	Places     []*Location
}

func NewWorld(plyr *Player, places []*Location) *World {
	return &World{
		PlayerData: plyr,
		Places:     places,
	}
}
