package tic80

import "unsafe"

const (
	SCREEN_TILE_HEIGHT  int32 = 17
	SCREEN_TILE_WIDTH         = 30
	SCREEN_PIXEL_HEIGHT       = 136
	SCREEN_PIXEL_WIDTH        = 240
)
const (
	MAP_TILE_HEIGHT int32 = 136
	MAP_TILE_WIDTH  int32 = 240
)

func Map(x, y int32) {
	var safeX, safeY int32

	if x >= 0 && x <= (MAP_TILE_WIDTH-SCREEN_TILE_WIDTH) {
		safeX = x
	} else if x < 0 {
		safeX = 0
	} else {
		safeX = (MAP_TILE_WIDTH - SCREEN_TILE_WIDTH)
	}

	if y >= 0 && y <= (MAP_TILE_HEIGHT-SCREEN_TILE_HEIGHT) {
		safeY = y
	} else if y < 0 {
		safeY = 0
	} else {
		safeY = (MAP_TILE_HEIGHT - SCREEN_TILE_HEIGHT)
	}

	_map(
		safeX, safeY,
		// Full screen of tiles
		SCREEN_TILE_WIDTH, SCREEN_TILE_HEIGHT,
		// Alway from the top left corner
		0, 0,
		// Transparency
		unsafe.Pointer(nil), 0,
		// Scale
		1,
		// Unused
		0,
	)
}
