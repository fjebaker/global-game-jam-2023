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

func DebugUpdate(world *World, player *Player) {
	xs := strconv.Itoa(int(player.X))
	ys := strconv.Itoa(int(player.Y))

	xw, yw, xo, yo := worldToTile(player.X, player.Y)

	location := "Player: x:" + xs + " y:" + ys + " speed:" + strconv.Itoa(int(player.Speed))
	world_location := "Tile: x:" + strconv.Itoa(int(xw)) + " y:" + strconv.Itoa(int(yw))
	offset_location := "Offset: x:" + strconv.Itoa(int(xo)) + " y:" + strconv.Itoa(int(yo))
	tile := world.GetMapTile(player.X, player.Y)
	tile_type := "Tile Type: " + strconv.Itoa(int(tile))

	x, y := player.GetInfront()
	xt, yt, _, _ := worldToTile(x, y)
	in_tile := world.GetMapTile(x, y)
	in_tile_location := "Infront Tile Coord: x:" + strconv.Itoa(int(xt)) + " y:" + strconv.Itoa(int(yt))
	in_tile_type := "Infront Type: " + strconv.Itoa(int(in_tile))

	tic80.Print(&location, 1, 1, 15, true, 1, false)
	tic80.Print(&world_location, 1, 8, 15, true, 1, false)
	tic80.Print(&offset_location, 1, 14, 15, true, 1, false)
	tic80.Print(&tile_type, 1, 21, 15, true, 1, false)
	tic80.Print(&in_tile_location, 1, 28, 15, true, 1, false)
	tic80.Print(&in_tile_type, 1, 35, 15, true, 1, false)

	has_message := "Has Item:"
	if player.HasItem {
		has_message = has_message + "True"
	} else {
		has_message = has_message + "False"
	}
	tic80.Print(&has_message, 1, 42, 15, true, 1, false)

	tic80.PaintPixel(x-player.X+PLAYER_OFFSET_X, y-player.Y+PLAYER_OFFSET_Y, 11)
}
