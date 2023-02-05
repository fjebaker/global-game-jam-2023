package cart

import (
	"cart/tic80"
	"math/rand"
)

const (
	OVERFLOW_MODULO_TIME int32 = 3600
)

var (
	_mouse        tic80.MouseData
	_t            int32
	_player       Player
	_rabbit       Rabbit
	_world        World
	_game         Game
	_desired_item RetrievableItem
)

func Start() {
	_t = 0
	_game = NewGame()
	_desired_item = NewRetrievableItem()
	_player = NewPlayer(PLAYER_START_POSITION_X, PLAYER_START_POSITION_Y)
	_rabbit = NewRabbit(100, 50, RABBIT_START_POSITION_X, RABBIT_START_POSITION_Y)
	_world = NewWorld(&_player)

	tic80.Music(GAME_PLAY_MUSIC, -1, -1, true, false, -1, -1)
}

// mainloop
func Loop() {
	tic80.Clear(0)
	_mouse.Update()

	if _game.State != GAME_STATE_OVER {
		_player.HandleInteraction(_t)

		_desired_item.Update(_t, &_game, &_player, &_rabbit)
		if _game.State != GAME_STATE_WIN {
			_player.Update(_t, &_game, &_world, &_desired_item, &_rabbit)
		}
		_rabbit.Update(_t, &_game, &_player)
		_world.Update(_t, &_game, &_player)
	}

	_world.Draw()
	_rabbit.Draw()
	if _game.State != GAME_STATE_WIN {
		_player.Draw()
		_desired_item.Draw()
	}

	_t = _t + 1
	// avoid overflows
	_t = _t % OVERFLOW_MODULO_TIME
}

// utility methods

func TimeSince(t, t_start int32) int32 {
	if t < t_start {
		return (t + OVERFLOW_MODULO_TIME) - t_start
	} else {
		return t - t_start
	}
}

func RandInt(min, max int) int {
	return min + rand.Intn(max-min)
}
