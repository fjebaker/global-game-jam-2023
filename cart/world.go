package cart

import "cart/tic80"

const (
	WORLD_LEFT_X int32 = 0
	// The "right edge" of the world occurs when the left edge
	// is one screen's worth of tiles away (because of how we draw)
	WORLD_RIGHT_X  = tic80.MAP_MAX_X - tic80.SCREEN_TILE_WIDTH*2
	WORLD_GROUND_Y = 13
	WORLD_BOTTOM_Y = tic80.MAP_MAX_Y

	WORLD_BACKGROUND_X = tic80.MAP_MAX_X - tic80.SCREEN_TILE_WIDTH
)

type World struct {
	X, Y             int32
	OffsetX, OffsetY int32
}

///////////////////////////////////////////////////////////////////////////////
// Methods

func NewWorld(player *Player) World {
	tileX, tileY, offsetX, offsetY := worldToTile(player.X, player.Y)

	return World{tileX, tileY, offsetX, offsetY}
}

func (world *World) Draw(t int32) {
	// The background is the entire right screenwidth of the map tiles
	tic80.Map(WORLD_BACKGROUND_X, world.Y, 0, world.OffsetY)
	// World.X, World.Y is the tile coordinates of the upper left corner
	tic80.Map(world.X, world.Y, world.OffsetX, world.OffsetY)
}

func (world *World) IsInBounds(x, y int32) bool {
	tileX, tileY, _, _ := worldToTile(x-PLAYER_OFFSET_X, y)

	if tileX <= WORLD_LEFT_X || WORLD_RIGHT_X <= tileX {
		return false
	}
	if y <= PLAYER_OFFSET_Y || WORLD_BOTTOM_Y <= tileY {
		return false
	}

	return true
}

func (world *World) Update(t int32, player *Player) {
	tileX, tileY, offsetX, offsetY := worldToTile(
		player.X-PLAYER_OFFSET_X,
		player.Y-PLAYER_OFFSET_Y,
	)

	world.X = tileX
	world.Y = tileY
	world.OffsetX = offsetX
	world.OffsetY = offsetY
}

///////////////////////////////////////////////////////////////////////////////
// Utils

func worldToTile(x, y int32) (int32, int32, int32, int32) {
	tileX := x / 8
	tileY := y / 8
	offsetX := x % 8
	offsetY := y % 8

	return tileX, tileY, offsetX, offsetY
}
