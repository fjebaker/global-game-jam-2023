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
	RABBIT_WIN_DURATION    = 120

	RABBIT_THANKS_SOUND   = 60
	RABBIT_SOUND_CHANNEL  = 2
	RABBIT_SOUND_DURATION = (3 * 60)

	RABBIT_STARTING_HEALTH = 120 // in seconds
	RABBIT_HURT_HEALTH     = 60  // at which point new frames used
	RABBIT_STARVING_FACTOR = 10  // in frames; multiplied by tree's life bar
	RABBIT_HEALING_AMOUNT  = 20

	RABBIT_HEALINGS_WIN = 10 // Number of healings to win the game
)

type Rabbit struct {
	X, Y           int32
	MapX, MapY     int32
	Frame          int32
	Sprite         tic80.Sprite
	HappySfx       tic80.SoundEffect
	Heart          TimedSprite
	ShowHeart      bool
	ShowWin        bool
	Health         int32
	HealsRemaining int32
	DeathClock     int32
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
		false,
		RABBIT_STARTING_HEALTH,
		RABBIT_HEALINGS_WIN,
		0,
	}
}

func (rabbit *Rabbit) Draw() {
	rabbit.Sprite.Draw(rabbit.X, rabbit.Y)
	rabbit.Heart.Draw(rabbit.X+6, rabbit.Y+6)
}

func (rabbit *Rabbit) IsDead() bool {
	return rabbit.Health == 0
}

func (rabbit *Rabbit) Heal(game *Game) {
	rabbit.HappySfx.Play()
	rabbit.ShowHeart = true
	rabbit.Health = rabbit.Health + RABBIT_HEALING_AMOUNT

	if rabbit.HealsRemaining > 0 {
		rabbit.HealsRemaining -= 1
		if rabbit.HealsRemaining == 0 {
			game.ChangeState(GAME_STATE_WIN)
			rabbit.ShowHeart = false
			rabbit.ShowWin = true
		}
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

func (rabbit *Rabbit) Update(t int32, game *Game, player *Player) {
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
	} else if rabbit.ShowWin {
		rabbit.Heart.WithZoom = true
		rabbit.Heart.Duration = RABBIT_WIN_DURATION
		rabbit.Heart.StartShowing(t)
		rabbit.ShowWin = false
	}

	// Run the animations
	rabbit.Heart.Update(t)
	if t%45 == 0 {
		rabbit.switchIdleFrame()
	}

	// deal with slow death
	rabbit.DieALittle(t, int32(game.TreeLife))
	// if we died this frame, load dead assets
	if rabbit.IsDead() {
		rabbit.Sprite.Id = RABBIT_DEAD_FRAME
	}

	// Deal with game win
	if game.State == GAME_STATE_WIN && !rabbit.Heart.Show {
		// When we've won, and the heart has stopped animating, go on
	}
}

func (rabbit *Rabbit) DieALittle(t int32, treeLife int32) {
	// every second decrease health and bump DeathClock
	if TimeSince(t, rabbit.DeathClock) >= (RABBIT_STARVING_FACTOR * treeLife) {
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
	if rabbit.Health <= RABBIT_HURT_HEALTH {
		rabbit.Sprite.Id = RABBIT_HURT_FRAME + (4 * rabbit.Frame)
	} else {
		rabbit.Sprite.Id = RABBIT_MAIN_FRAME + (4 * rabbit.Frame)
	}
}
