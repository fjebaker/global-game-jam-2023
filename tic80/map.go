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
)

func Map(tileX, tileY, offsetX, offsetY int32) {
	var safeX, safeY int32

	if tileX >= 0 && tileX <= (MAP_MAX_X-SCREEN_TILE_WIDTH) {
		safeX = tileX
	} else if tileX < 0 {
		safeX = 0
	} else {
		safeX = MAP_MAX_X - SCREEN_TILE_WIDTH
	}

	if tileY >= 0 && tileY <= (MAP_MAX_Y-SCREEN_TILE_HEIGHT) {
		safeY = tileY
	} else if tileY < 0 {
		safeY = 0
	} else {
		safeY = MAP_MAX_Y - SCREEN_TILE_HEIGHT
	}

	_map(
		safeX, safeY,
		// Full screen of tiles
		SCREEN_TILE_WIDTH+1, SCREEN_TILE_HEIGHT+1,
		// Alway from the top left corner
		-offsetX, -offsetY,
		// Transparency
		unsafe.Pointer(nil), 0,
		// Scale
		1,
		// Unused
		0,
	)
}
