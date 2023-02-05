package cart

import "cart/tic80"

type GameState int32

const (
	GAME_STATE_PLAYING GameState = iota
	GAME_STATE_OVER
)

const (
	GAME_DEATH_SOUND    = 59
	GAME_DEATH_CHANNEL  = 3
	GAME_DEATH_DURATION = 30

	GAME_PLAY_MUSIC = 0
	GAME_OVER_MUSIC = 1
)

type Game struct {
	State    GameState
	DeathSfx tic80.SoundEffect
}

///////////////////////////////////////////////////////////////////////////////
// Methods

func NewGame() Game {
	sfx := tic80.NewSoundEffect(
		GAME_DEATH_SOUND, GAME_DEATH_CHANNEL, GAME_DEATH_DURATION,
	)

	return Game{GAME_STATE_PLAYING, sfx}
}

func (game *Game) ChangeState(state GameState) {
	game.State = state

	if state == GAME_STATE_OVER {
		game.DeathSfx.Play()
		// Play a sad song
		tic80.Music(GAME_OVER_MUSIC, -1, -1, true, false, -1, -1)
	}
}
