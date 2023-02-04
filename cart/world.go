package cart

import "cart/tic80"

type World struct {
	X, Y int32
	OffsetX, OffsetY int32
}

func NewWorld(player *Player) World {
	tileX, tileY := worldPosition(player)
	return World{tileX, tileY, 0, 0}
}

func worldPosition(player *Player) (int32, int32) {
	playerTileX := (player.MapX / 8) - (tic80.SCREEN_TILE_WIDTH / 2)
	playerTileY := (player.MapY / 8) - (tic80.SCREEN_TILE_HEIGHT / 2)
	return playerTileX, playerTileY
}

func (world *World) Draw(t int32) {
	tic80.Map(world.X, world.Y, world.OffsetX, world.OffsetY)
}

func (world *World) Update(t int32, player *Player) {
	tileX, tileY := worldPosition(player)
	world.X = tileX
	world.Y = tileY
	world.OffsetX = player.MapX % 8
	world.OffsetY = player.MapY % 8
}
