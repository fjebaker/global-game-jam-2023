package cart

import "cart/tic80"

type GameState int32

const (
	GAME_STATE_PLAYING GameState = iota
	GAME_STATE_OVER
)

type DesireableItem int32

const (
	ITEM_GREEN_MUSHROOM DesireableItem = 32
	ITEM_RADISH         DesireableItem = 33
	ITEM_COIN           DesireableItem = 34
	ITEM_BLUE_MUSHROOM  DesireableItem = 35
	ITEM_RED_MUSHROOM   DesireableItem = 48
	ITEM_CLAY_POT       DesireableItem = 49
	ITEM_RUBY_RING      DesireableItem = 50
	ITEM_RED_THING      DesireableItem = 100
	ITEM_GREEN_THING    DesireableItem = 101
)

const DESIREABLE_ITEM_NUMBER = 9

type Game struct {
	State       GameState
	DesiredItem DesireableItem
}

///////////////////////////////////////////////////////////////////////////////
// Methods

func NewGame() Game {
	return Game{GAME_STATE_PLAYING, ITEM_RED_MUSHROOM}
}

func (game *Game) ChangeState(state GameState) {
	game.State = state
	if state == GAME_STATE_OVER {
		sfx := tic80.NewSoundEffect(59, 3, 30)
		sfx.Play()
		tic80.Music(1, -1, -1, true, false, -1, -1)
	}
}

func (game *Game) NewDesiredItem() {
	game.DesiredItem = newDesiredItem()
}

// utility

func newDesiredItem() DesireableItem {
	i := RandInt(0, DESIREABLE_ITEM_NUMBER)
	switch i {
	case 0:
		return ITEM_GREEN_MUSHROOM
	case 1:
		return ITEM_RADISH
	case 2:
		return ITEM_COIN
	case 3:
		return ITEM_BLUE_MUSHROOM
	case 4:
		return ITEM_RED_MUSHROOM
	case 5:
		return ITEM_CLAY_POT
	case 6:
		return ITEM_RUBY_RING
	case 7:
		return ITEM_RED_THING
	case 8:
		return ITEM_GREEN_THING
	// just incase
	default:
		return ITEM_RADISH
	}
}
