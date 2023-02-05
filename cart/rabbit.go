package cart

import "cart/tic80"

type Rabbit struct {
	X, Y       int32
	MapX, MapY int32
	Frame      int32
	Sprite     tic80.Sprite
	HappySfx   tic80.SoundEffect
	Heart      TimedSprite
	ShowHeart  bool
}

const rabbit_main_frame = 272

const (
	RABBIT_START_POSITION_X = (104 * 8)
	RABBIT_START_POSITION_Y = (12 * 8)

	RABBIT_DETECTION_ZONE = 30
)

func NewRabbit(x, y, mapx, mapy int32) Rabbit {
	sprite := tic80.SquareSprite(rabbit_main_frame, 4)
	sfx := tic80.NewSoundEffect(60, 2)
	sfx.Duration = 180
	return Rabbit{
		x, y,
		mapx, mapy,
		0,
		sprite,
		sfx,
		NewTimedSprite(
			tic80.SquareSprite(352, 2),
			30,
		),
		false,
	}
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
	rabbit.Sprite.Draw(rabbit.X, rabbit.Y)
	rabbit.Heart.Draw(t, rabbit.X+6, rabbit.Y+6)
}

func (rabbit *Rabbit) Update(t int32, player *Player, game *Game) {
	rabbit.X = rabbit.MapX - (player.X) + PLAYER_OFFSET_X
	rabbit.Y = rabbit.MapY - (player.Y) + PLAYER_OFFSET_Y

	if rabbit.ShowHeart {
		rabbit.Heart.StartShowing(t)
		rabbit.ShowHeart = false
	}
}

func (rabbit *Rabbit) PointInZone(x, y int32) bool {
	// calculate the distance
	dx := x - rabbit.MapX
	// add a little padding left of the rabbit too
	x_condition := dx >= -2 && dx <= RABBIT_DETECTION_ZONE
	y_condition := y-rabbit.MapY <= RABBIT_DETECTION_ZONE
	return x_condition && y_condition
}
