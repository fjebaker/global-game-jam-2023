package cart

import (
	"cart/tic80"
	"strconv"
)

var (
	_is_enabled bool
)

func DebugStart() {
	_is_enabled = false
}

func DebugUpdate(player *Player) {
	xs := strconv.Itoa(int(player.X))
	ys := strconv.Itoa(int(player.Y))

	xw, yw, xo, yo := worldToTile(player.X, player.Y)

	location := "Player: x:" + xs + " y:" + ys
	world_location := "Tile: x:" + strconv.Itoa(int(xw)) + " y:" + strconv.Itoa(int(yw))
	offset_location := "Offset: x:" + strconv.Itoa(int(xo)) + " y:" + strconv.Itoa(int(yo))
	tic80.Print(&location, 1, 1, 15, true, 1, false)
	tic80.Print(&world_location, 1, 8, 15, true, 1, false)
	tic80.Print(&offset_location, 1, 15, 15, true, 1, false)

	tic80.PaintPixel(player.X-PLAYER_OFFSET_X, player.Y-PLAYER_OFFSET_Y, 10)
}
