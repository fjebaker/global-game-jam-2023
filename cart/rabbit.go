package cart

import "cart/tic80"

const (
	RABBIT_START_POSITION_X = (104 * 8)
	RABBIT_START_POSITION_Y = (12 * 8)

	RABBIT_DETECTION_ZONE = 30

	RABBIT_MAIN_FRAME      = 272
	RABBIT_HURT_FRAME      = 384
	RABBIT_DEAD_FRAME      = 280
	RABBIT_FRAME_SIZE      = 4
	RABBIT_THANKS_FRAME    = 352
	RABBIT_THANKS_SIZE     = 2
	RABBIT_THANKS_DURATION = 30

	RABBIT_THANKS_SOUND   = 60
	RABBIT_SOUND_CHANNEL  = 2
	RABBIT_SOUND_DURATION = (3 * 60)

	RABBIT_STARTING_HEALTH = 120  // in seconds
	RABBIT_STARVING_RATE   = 60 // in frames
)

type Rabbit struct {
	X, Y       int32
	MapX, MapY int32
	Frame      int32
	Sprite     tic80.Sprite
	HappySfx   tic80.SoundEffect
	Heart      TimedSprite
	ShowHeart  bool
	Health     int32
	DeathClock int32
}

///////////////////////////////////////////////////////////////////////////////
// Methods

func NewRabbit(x, y, mapx, mapy int32) Rabbit {
	sprite := tic80.SquareSprite(RABBIT_MAIN_FRAME, RABBIT_FRAME_SIZE)
	sfx := tic80.NewSoundEffect(
		RABBIT_THANKS_SOUND, RABBIT_SOUND_CHANNEL, RABBIT_SOUND_DURATION,
	)
	return Rabbit{
		x, y,
		mapx, mapy,
		0,
		sprite,
		sfx,
		NewTimedSprite(
			tic80.SquareSprite(RABBIT_THANKS_FRAME, RABBIT_THANKS_SIZE),
			RABBIT_THANKS_DURATION,
		),
		false,
		RABBIT_STARTING_HEALTH,
		0,
	}
}

func (rabbit *Rabbit) Draw() {
	rabbit.Sprite.Draw(rabbit.X, rabbit.Y)
	rabbit.Heart.Draw(rabbit.X+6, rabbit.Y+6)
}

func (rabbit *Rabbit) PointInZone(x, y int32) bool {
	// calculate the distance
	dx := x - rabbit.MapX
	// add a little padding left of the rabbit too
	x_condition := dx >= -2 && dx <= RABBIT_DETECTION_ZONE
	y_condition := y-rabbit.MapY <= RABBIT_DETECTION_ZONE
	return x_condition && y_condition
}

func (rabbit *Rabbit) Update(t int32, player *Player, game *Game) {
	rabbit.X = rabbit.MapX - (player.X) + PLAYER_OFFSET_X
	rabbit.Y = rabbit.MapY - (player.Y) + PLAYER_OFFSET_Y

	// if health is 0 we don't need to update anything other than position
	// cus the rabbit is dead
	if rabbit.IsDead() {
		return
	}

	if rabbit.ShowHeart {
		rabbit.Heart.StartShowing(t)
		rabbit.ShowHeart = false
	}

	// Run the animations
	rabbit.Heart.Update(t)
	if t%45 == 0 {
		rabbit.switchIdleFrame()
	}

	// deal with slow death
	rabbit.DieALittle(t)
	// if we died this frame, load dead assets
	if rabbit.IsDead() {
		rabbit.Sprite.Id = RABBIT_DEAD_FRAME
	}
}

func (rabbit *Rabbit) DieALittle(t int32) {
	// every second decrease health and bump DeathClock
	if TimeSince(t, rabbit.DeathClock) >= RABBIT_STARVING_RATE {
		rabbit.DeathClock = t
		rabbit.Health = rabbit.Health - 1
	}
}

///////////////////////////////////////////////////////////////////////////////
// Utility

func (rabbit *Rabbit) switchIdleFrame() {
	if rabbit.Frame == 0 {
		rabbit.Frame = 1
	} else {
		rabbit.Frame = 0
	}
	rabbit.Sprite.Id = RABBIT_MAIN_FRAME + (4 * rabbit.Frame)
}

func (rabbit *Rabbit) IsDead() bool {
	return rabbit.Health == 0
}
