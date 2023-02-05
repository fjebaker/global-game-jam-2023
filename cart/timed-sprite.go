package cart

import "cart/tic80"

type TimedSprite struct {
	StartTime int32
	Duration  int32
	Sprite    tic80.Sprite
	Show      bool
}

func NewTimedSprite(sprite tic80.Sprite, duration int32) TimedSprite {
	return TimedSprite{0, duration, sprite, false}
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
		if TimeSince(t, sprite.StartTime) > sprite.Duration {
			sprite.Show = false
		}
	}
}
