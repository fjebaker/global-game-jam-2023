package cart

import "cart/tic80"

type GameState int32

const (
	GAME_STATE_PLAYING GameState = iota
	GAME_STATE_OVER
)

type Game struct {
	State GameState
}

///////////////////////////////////////////////////////////////////////////////
// Methods

func NewGame() Game {
	return Game{GAME_STATE_PLAYING}
}

func (game *Game) ChangeState(state GameState) {
	game.State = state
	if state == GAME_STATE_OVER {
		sfx := tic80.NewSoundEffect(59, 3)
		sfx.Play()
		tic80.Music(1, -1, -1, true, false, -1, -1)
	}
}
