package cart

import (
	"cart/tic80"
)

const (
	OVERFLOW_MODULO_TIME int32 = 3600
)

var (
	_mouse  tic80.MouseData
	_t      int32
	_player Player
	_rabbit Rabbit
	_world  World
)

func Start() {
	_t = 0
	_player = NewPlayer(PLAYER_START_POSITION_X, PLAYER_START_POSITION_Y)
	_rabbit = NewRabbit(100, 50, RABBIT_START_POSITION_X, RABBIT_START_POSITION_Y)
	_world = NewWorld(&_player)
}

// mainloop
func Loop() {
	tic80.Clear(0)
	_mouse.Update()

	_player.HandleInteraction(_t)
	_player.Update(_t, &_world)
	_rabbit.Update(_t, &_player)
	_world.Update(_t, &_player)

	_world.Draw(_t)
	_rabbit.Draw(_t)
	_player.Draw(_t)

	_t = _t + 1
	// avoid overflows
	_t = _t % OVERFLOW_MODULO_TIME
}
