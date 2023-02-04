package cart

import "cart/tic80"

const (
	PLAYER_DELTA_X = 4
	PLAYER_DELTA_Y = 4

	// Keep the player centered on screen
	PLAYER_OFFSET_X int32 = 120 - PLAYER_DELTA_X
	PLAYER_OFFSET_Y       = 67 - PLAYER_DELTA_Y

	PLAYER_START_POSITION_X = 97 * 8
	PLAYER_START_POSITION_Y = 14 * 8
)

type Player struct {
	X, Y    int32
	Frame   int32
	Sprite  tic80.Sprite
	Move_fx tic80.SoundEffect
	Speed   int32
	Moving  bool
}

func NewPlayer(worldX, worldY int32) Player {
	sprite := tic80.SquareSprite(258, 1)
	sprite.Rotate = tic80.ROTATE_RIGHT
	sfx := tic80.NewSoundEffect(61, 0)

	return Player{worldX, worldY, 0, sprite, sfx, 10, false}
}

const player_main_frame = 256

func (player *Player) incrementFrame() {
	player.Frame = (player.Frame + 1) % 5
	player.Sprite.Id = player_main_frame + player.Frame
}

func (player *Player) Draw(t int32) {
	var mod int32
	if player.Moving {
		mod = 5
	} else {
		mod = 12
	}
	if t%mod == 0 {
		player.incrementFrame()
	}
	player.Sprite.Draw(PLAYER_OFFSET_X, PLAYER_OFFSET_Y)
}

func (player *Player) HandleInteraction(t int32) {
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

func (player *Player) move(world *World) {
	maybeX := player.X
	maybeY := player.Y

	switch player.Sprite.Rotate {
	case tic80.ROTATE_NONE:
		maybeY -= 1
	case tic80.ROTATE_DOWN:
		maybeY += 1
	case tic80.ROTATE_RIGHT:
		maybeX += 1
	case tic80.ROTATE_LEFT:
		maybeX -= 1
	}

	if world.IsInBounds(maybeX, maybeY) {
		player.X = maybeX
		player.Y = maybeY
	}
}

func (player *Player) Update(t int32, world *World) {
	if player.Moving {
		// check sfx update
		if player.Move_fx.IsPlaying(t) == false {
			player.Move_fx.PlayRecordTime(t)
		}
		// check whether to advance location
		if (t*player.Speed)%30 == 0 {
			player.move(world)
		}
	}
}
