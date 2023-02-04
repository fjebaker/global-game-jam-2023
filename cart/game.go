package cart

type GameState int32

const (
	GAME_STATE_PLAYING GameState = iota
)

type Game struct {
	State GameState
}

///////////////////////////////////////////////////////////////////////////////
// Methods

func NewGame() Game {
	return Game{GAME_STATE_PLAYING}
}

func (game *Game) Update(t int32) {
}
