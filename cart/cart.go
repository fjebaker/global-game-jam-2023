package cart

import (
	"cart/tic80"
)

var (
	_mouse  tic80.MouseData
	_t      int32
	_player Player
	_world  World
)

func Start() {
	_t = 0
	_player = NewPlayer(PLAYER_START_POSITION_X, PLAYER_START_POSITION_Y)
	_world = NewWorld(&_player)
}

// mainloop
func Loop() {
	tic80.Clear(0)
	_mouse.Update()

	_player.HandleInteraction(_t)
	_player.Update(_t, &_world)
	_world.Update(_t, &_player)

	_world.Draw(_t)
	_player.Draw(_t)

	_t = _t + 1
	// avoid overflows
	_t = _t % 3600
}
