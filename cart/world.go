package cart

import "cart/tic80"

type World struct {
	X, Y int32
}

func NewWorld(player *Player) World {
	tileX, tileY := worldPosition(player)
	return World{tileX, tileY}
}

func worldPosition(player *Player) (int32, int32) {
	playerTileX := (player.X / 8) - (tic80.SCREEN_TILE_WIDTH / 2)
	playerTileY := (player.Y / 8) - (tic80.SCREEN_TILE_HEIGHT / 2)
	return playerTileX, playerTileY
}

func (world *World) Draw(t int32) {
	tic80.Map(world.X, world.Y)
}

func (world *World) Update(t int32, player *Player) {
	tileX, tileY := worldPosition(player)
	world.X = tileX
	world.Y = tileY
}
