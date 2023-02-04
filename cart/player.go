package cart

import "cart/tic80"

const (
	PLAYER_DELTA_X = 4
	PLAYER_DELTA_Y = 4

	// Keep the player centered on screen
	PLAYER_OFFSET_X int32 = 120 - PLAYER_DELTA_X
	PLAYER_OFFSET_Y       = 67 - PLAYER_DELTA_Y

	PLAYER_START_POSITION_X = 97*8 + 4*8
	PLAYER_START_POSITION_Y = 14 * 8
)

type Player struct {
	X, Y       int32
	Frame      int32
	Sprite     tic80.Sprite
	Move_fx    tic80.SoundEffect
	Speed      int32
	Dead       bool
	Digging    bool
	Moving     bool
	HasItem    bool
	ItemSprite *tic80.Sprite
}

func NewPlayer(worldX, worldY int32, desired_item_sprite *tic80.Sprite) Player {
	sprite := tic80.SquareSprite(258, 1)
	sprite.Rotate = tic80.ROTATE_RIGHT
	sfx := tic80.NewSoundEffect(61, 3)

	return Player{worldX, worldY, 0, sprite, sfx, 4, false, false, false, false, desired_item_sprite}
}

const (
	PLAYER_MAIN_FRAME = 256
	// DEBUG FRAME
	// PLAYER_MAIN_FRAME = 336
	PLAYER_DEAD_FRAME = 261
)

///////////////////////////////////////////////////////////////////////////////
// Methods

func (player *Player) Draw(t int32) {
	var mod int32
	if player.Moving {
		if player.Speed == 1 {
			mod = 2
		} else {
			mod = 5
		}
	} else {
		mod = 12
	}
	if t%mod == 0 {
		player.incrementFrame()
	}
	player.Sprite.Draw(PLAYER_OFFSET_X, PLAYER_OFFSET_Y)

	player.DrawTooltips()
}

func (player *Player) GetInfront() (int32, int32) {
	var x, y int32
	switch player.Sprite.Rotate {
	case tic80.ROTATE_NONE:
		y = player.Y - 8 + PLAYER_DELTA_Y - 1
		x = player.X
	case tic80.ROTATE_DOWN:
		y = player.Y + 8 - PLAYER_DELTA_Y
		x = player.X
	case tic80.ROTATE_RIGHT:
		y = player.Y
		x = player.X + 8 - PLAYER_DELTA_X
	case tic80.ROTATE_LEFT:
		y = player.Y
		x = player.X - 8 + PLAYER_DELTA_Y - 1
	}

	return x + PLAYER_DELTA_X, y + PLAYER_DELTA_Y
}

func (player *Player) HandleInteraction(t int32) {
	// We always check for digging
	player.Digging = tic80.BUTTON_B.IsPressed()

	if tic80.BUTTON_A.IsPressed() {
		player.Speed = 1
		player.Move_fx.Note = 2
	} else {
		player.Speed = 3
		player.Move_fx.Note = 0
	}

	if tic80.BUTTON_UP.IsPressed() {
		player.Sprite.Rotate = tic80.ROTATE_NONE
		player.Moving = true
		return
	}
	if tic80.BUTTON_RIGHT.IsPressed() {
		player.Sprite.Rotate = tic80.ROTATE_RIGHT
		player.Moving = true
		return
	}
	if tic80.BUTTON_LEFT.IsPressed() {
		player.Sprite.Rotate = tic80.ROTATE_LEFT
		player.Moving = true
		return
	}
	if tic80.BUTTON_DOWN.IsPressed() {
		player.Sprite.Rotate = tic80.ROTATE_DOWN
		player.Moving = true
		return
	}
	player.Moving = false
	player.Move_fx.Stop()
}

func (player *Player) Update(t int32, world *World, game *Game) {
	if player.Moving {
		// check sfx update
		if player.Move_fx.IsPlaying(t, OVERFLOW_MODULO_TIME) == false {
			player.Move_fx.PlayRecordTime(t)
		}
		// check whether to advance location
		if ((t*15)/10)%player.Speed == 0 {
			player.move(world)
		}
	}

	// check what is infront
	x, y := player.GetInfront()
	tileIndex := world.GetMapTile(x, y)

	if world.IsDeadly(tileIndex) {
		game.ChangeState(GAME_STATE_OVER)
		player.Dead = true
	}

	if player.Digging {
		switch {
		case world.IsDirt(tileIndex):
			world.DigTile(x, y)
		case world.IsItem(tileIndex):
			world.CollectItem(x, y)
		}
		// check if the item was what we wanted
		if tileIndex == int32(game.DesiredItem) {
			player.HasItem = true
		}
	}
}

///////////////////////////////////////////////////////////////////////////////
// Utils

func (player *Player) incrementFrame() {
	if !player.Dead {
		player.Frame = (player.Frame + 1) % 5
		player.Sprite.Id = PLAYER_MAIN_FRAME + player.Frame
	} else {
		player.Sprite.Id = PLAYER_DEAD_FRAME
	}
}

func (player *Player) move(world *World) {
	x, y := player.GetInfront()
	// if we are trying to move out of bounds
	// don't
	if !world.IsInBounds(x, y) {
		return
	}
	// What does the tile in that position contain?
	tileIndex := world.GetMapTile(x, y)

	if world.IsIndestructible(tileIndex) || world.IsDirt(tileIndex) {
		return
	}

	switch player.Sprite.Rotate {
	case tic80.ROTATE_NONE:
		player.Y -= 1
	case tic80.ROTATE_DOWN:
		player.Y += 1
	case tic80.ROTATE_RIGHT:
		player.X += 1
	case tic80.ROTATE_LEFT:
		player.X -= 1
	}
}

func (player *Player) DrawTooltips() {
	tic80.RectangleWithBorder(0, 0, 12, 12, 12, 9)
	if player.HasItem {
		player.ItemSprite.Draw(2, 2)
	}
}
