package cart

import "cart/tic80"

const (
	// The "left edge" of the world occurs in the middle of the
	// left-most screen because that's the _player's_ position limit
	WORLD_LEFT_X int32 = 15
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

func (world *World) CollectItem(x, y int32) {
	tileX, tileY, _, _ := worldToTile(x, y)
	tic80.MSet(tileX, tileY, tic80.MAP_EMPTY_TILE)

	// NOTE: This only works for items up to 2x2 tiles
	for dy := int32(-1); dy < 2; dy++ {
		for dx := int32(-1); dx < 2; dx++ {
			index := tic80.MGet(tileX+dx, tileY+dy)
			if tic80.FGet(index, tic80.MAP_TILE_ITEM_FLAG) {
				tic80.MSet(tileX+dx, tileY+dy, tic80.MAP_EMPTY_TILE)
			}
		}
	}

}

func (world *World) DigTile(x, y int32) {
	tileX, tileY, _, _ := worldToTile(x, y)
	tic80.MSet(tileX, tileY, tic80.MAP_EMPTY_TILE)
}

func (world *World) Draw() {
	// The background is the entire right screenwidth of the map tiles
	tic80.Map(WORLD_BACKGROUND_X, world.Y, 0, world.OffsetY)
	// World.X, World.Y is the tile coordinates of the upper left corner
	tic80.Map(world.X, world.Y, world.OffsetX, world.OffsetY)
}

func (world *World) GetMapTile(x, y int32) int32 {
	tileX, tileY, _, _ := worldToTile(x, y)
	return tic80.MGet(tileX, tileY)
}

func (world *World) IsDeadly(index int32) bool {
	return tic80.FGet(index, tic80.MAP_TILE_DEADLY_FLAG)
}

func (world *World) IsDirt(index int32) bool {
	return tic80.FGet(index, tic80.MAP_TILE_DIRT_FLAG)
}

func (world *World) IsIndestructible(index int32) bool {
	return tic80.FGet(index, tic80.MAP_TILE_INDESTRUCTIBLE_FLAG)
}

func (world *World) IsItem(index int32) bool {
	return tic80.FGet(index, tic80.MAP_TILE_ITEM_FLAG)
}

func (world *World) IsInBounds(x, y int32) bool {
	tileX, tileY, _, _ := worldToTile(x, y)

	if tileX <= WORLD_LEFT_X || WORLD_RIGHT_X <= tileX {
		return false
	}
	if tileY <= WORLD_GROUND_Y || WORLD_BOTTOM_Y <= tileY {
		return false
	}

	return true
}

func (world *World) Update(t int32, player *Player, game *Game) {
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
	offsetX := 8 - (x % 8)
	offsetY := 8 - (y % 8)
	tileX := x / 8
	tileY := y / 8

	return tileX, tileY, offsetX, offsetY
}
