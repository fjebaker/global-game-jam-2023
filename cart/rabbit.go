package cart

import "cart/tic80"

type Rabbit struct {
	X, Y       int32
	MapX, MapY int32
	Frame      int32
	Sprite     tic80.Sprite
}

const rabbit_main_frame = 272

const (
	RABBIT_START_POSITION_X = (104 * 8)
	RABBIT_START_POSITION_Y = (12 * 8)
)

func NewRabbit(x, y, mapx, mapy int32) Rabbit {
	sprite := tic80.SquareSprite(rabbit_main_frame, 4)
	return Rabbit{x, y, mapx, mapy, 0, sprite}
}

func (rabbit *Rabbit) switchIdleFrame() {
	if rabbit.Frame == 0 {
		rabbit.Frame = 1
	} else {
		rabbit.Frame = 0
	}
	rabbit.Sprite.Id = rabbit_main_frame + (4 * rabbit.Frame)
}

func (rabbit *Rabbit) Draw(t int32) {
	if t%45 == 0 {
		rabbit.switchIdleFrame()
	}
	rabbit.Sprite.Draw(rabbit.X+PLAYER_OFFSET_X, rabbit.Y+PLAYER_OFFSET_Y)
}

func (rabbit *Rabbit) Update(t int32, player *Player, game *Game) {
	rabbit.X = rabbit.MapX - (player.X)
	rabbit.Y = rabbit.MapY - (player.Y)
}
