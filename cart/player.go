package cart

import "cart/tic80"

type Player struct {
	X, Y    int32
	Frame   int32
	Sprite  tic80.Sprite
	Move_fx tic80.SoundEffect
	Speed   int32
	Moving  bool
}

func NewPlayer(x, y int32) Player {
	sprite := tic80.SquareSprite(258, 1)
	sprite.Rotate = F_UP
	sfx := tic80.NewSoundEffect(61)
	return Player{x, y, 0, sprite, 10, false}
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
	player.Sprite.Draw(player.X, player.Y)
}

func (player *Player) HandleInteraction(t int32) {
	if tic80.BUTTON_UP.IsPressed() {
		player.Sprite.Rotate = F_UP
		player.Moving = true
		return
	}
	if tic80.BUTTON_RIGHT.IsPressed() {
		player.Sprite.Rotate = F_RIGHT
		player.Moving = true
		return
	}
	if tic80.BUTTON_LEFT.IsPressed() {
		player.Sprite.Rotate = F_LEFT
		player.Moving = true
		return
	}
	if tic80.BUTTON_DOWN.IsPressed() {
		player.Sprite.Rotate = F_DOWN
		player.Moving = true
		return
	}
	player.Moving = false
}

func (player *Player) move() {
	switch player.Sprite.Rotate {
	case F_UP:
		player.Y = player.Y - 1
	case F_DOWN:
		player.Y = player.Y + 1
	case F_RIGHT:
		player.X = player.X + 1
	case F_LEFT:
		player.X = player.X - 1
	}
}

func (player *Player) Update(t int32) {
	if player.Moving && (t*player.Speed)%30 == 0 {
		player.move()
	}
}
