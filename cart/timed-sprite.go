package cart

import "cart/tic80"

const ZOOM_DELAY = 10

type TimedSprite struct {
	StartTime int32
	Duration  int32
	Sprite    tic80.Sprite
	Show      bool
	WithZoom  bool
}

func NewTimedSprite(sprite tic80.Sprite, duration int32) TimedSprite {
	return TimedSprite{0, duration, sprite, false, false}
}

func (sprite *TimedSprite) StartShowing(t int32) {
	sprite.StartTime = t
	sprite.Show = true
}

func (sprite *TimedSprite) Draw(x, y int32) {
	if sprite.Show {
		sprite.Sprite.Draw(x, y)
	}
}

func (sprite *TimedSprite) Update(t int32) {
	if sprite.Show {
		if sprite.WithZoom && TimeSince(t, sprite.StartTime)%10 == 0 {
			sprite.Sprite.Scale += 1
		}
		if TimeSince(t, sprite.StartTime) > sprite.Duration {
			sprite.Show = false
		}
	}
}
