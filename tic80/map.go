package tic80

import "unsafe"

const (
	SCREEN_TILE_HEIGHT  int32 = 17
	SCREEN_TILE_WIDTH         = 30
	SCREEN_PIXEL_HEIGHT       = 136
	SCREEN_PIXEL_WIDTH        = 240
)
const (
	MAP_MAX_X       int32 = 240
	MAP_MAX_Y             = 136
	MAP_TILE_HEIGHT       = 8
	MAP_TILE_WIDTH        = 8

	MAP_EMPTY_TILE = 255

	MAP_TILE_INDESTRUCTIBLE_FLAG = 0
	MAP_TILE_DIRT_FLAG           = 1
	MAP_TILE_ITEM_FLAG           = 2
	MAP_TILE_DEADLY_FLAG         = 4
)

func Map(tileX, tileY, offsetX, offsetY int32) {
	var safeX, safeY int32
	transparent := 0

	if tileX > 0 && tileX <= (MAP_MAX_X-SCREEN_TILE_WIDTH) {
		safeX = tileX
	} else if tileX < 0 {
		safeX = 0
	} else {
		safeX = MAP_MAX_X - SCREEN_TILE_WIDTH
	}

	if tileY > 0 && tileY <= (MAP_MAX_Y-SCREEN_TILE_HEIGHT) {
		safeY = tileY
	} else if tileY < 0 {
		safeY = 0
	} else {
		safeY = MAP_MAX_Y - SCREEN_TILE_HEIGHT
	}

	_map(
		safeX-1, safeY-1,
		// Full screen of tiles
		SCREEN_TILE_WIDTH+2, SCREEN_TILE_HEIGHT+2,
		// Alway from the top left corner
		offsetX-16,
		offsetY-16,
		// Transparency
		unsafe.Pointer(&transparent), 1,
		// Scale
		1,
		// Unused
		0,
	)
}

func MGet(tileX, tileY int32) int32 {
	if tileX < 0 || MAP_MAX_X < tileX+1 {
		return MAP_EMPTY_TILE
	}
	if tileY < 0 || MAP_MAX_Y < tileY+1 {
		return MAP_EMPTY_TILE
	}

	return _mget(tileX, tileY)
}

func MSet(tileX, tileY, value int32) {
	if 0 <= tileX && 0 <= tileY && tileX < MAP_MAX_X && tileY < MAP_MAX_Y {
		_mset(tileX, tileY, value)
	}
}
