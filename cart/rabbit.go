package cart

import "cart/tic80"

const (
	THOUGHT_BUBBLE_FADE  = 0
	THOUGHT_BUBBLE_SMALL = 1
	THOUGHT_BUBBLE_BIG   = 2
)

type Rabbit struct {
	X, Y        int32
	MapX, MapY  int32
	Frame       int32
	Sprite      tic80.Sprite
	ShowItem    bool
	BubbleState int32
	Ticker      int32
	ItemSprite  *tic80.Sprite
	HappySfx    tic80.SoundEffect
}

const rabbit_main_frame = 272
const (
	THOUGHT_BUBBLE_DELAY = 60
	THOUGHT_BUBBLE_ZONE  = 30
)

const (
	RABBIT_START_POSITION_X = (104 * 8)
	RABBIT_START_POSITION_Y = (12 * 8)
	RABBIT_ITEM_OFFSET      = 4
)

func NewRabbit(x, y, mapx, mapy int32, item_sprite *tic80.Sprite) Rabbit {
	sprite := tic80.SquareSprite(rabbit_main_frame, 4)
	sfx := tic80.NewSoundEffect(60, 2, 30)
	sfx.Duration = 180
	return Rabbit{
		x, y,
		mapx, mapy,
		0,
		sprite,
		false,
		THOUGHT_BUBBLE_SMALL,
		0,
		item_sprite,
		sfx,
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

func (rabbit *Rabbit) drawThoughtBubble(x, y, t int32) {
	x_item := x - RABBIT_ITEM_OFFSET
	y_item := y - RABBIT_ITEM_OFFSET

	if rabbit.BubbleState >= THOUGHT_BUBBLE_SMALL {
		tic80.EllipseWithBorder(x_item+10, y_item+11, 2, 1, 12, 13)
	}
	if rabbit.BubbleState >= THOUGHT_BUBBLE_BIG {
		tic80.EllipseWithBorder(x_item+3, y_item+3, 7, 6, 12, 13)
		if rabbit.ItemSprite != nil {
			rabbit.ItemSprite.Draw(x_item, y_item)
		}
	}
}

func (rabbit *Rabbit) Draw(t int32) {
	if t%45 == 0 {
		rabbit.switchIdleFrame()
	}

	x := rabbit.X + PLAYER_OFFSET_X
	y := rabbit.Y + PLAYER_OFFSET_Y
	rabbit.Sprite.Draw(x, y)

	if rabbit.ShowItem {
		rabbit.drawThoughtBubble(x, y, t)
	}
}

func (rabbit *Rabbit) Update(t int32, player *Player, game *Game) {
	rabbit.X = rabbit.MapX - (player.X)
	rabbit.Y = rabbit.MapY - (player.Y)

	if rabbit.PointInZone(player.X, player.Y) {
		// does the player currently have the item
		if player.HasItem {
			rabbit.HappySfx.Play()
			player.HasItem = false
			game.NewDesiredItem()
			rabbit.ItemSprite.Id = int32(game.DesiredItem)
			rabbit.SetShowItem(t, false)
		}
		rabbit.SetShowItem(t, true)
	} else {
		rabbit.SetShowItem(t, false)
	}

	// update the state of the thought bubble
	if rabbit.ShowItem {
		if rabbit.BubbleState == THOUGHT_BUBBLE_SMALL {
			if TimeSince(t, rabbit.Ticker) > THOUGHT_BUBBLE_DELAY {
				rabbit.BubbleState = THOUGHT_BUBBLE_BIG
			}
		} else if rabbit.BubbleState == THOUGHT_BUBBLE_FADE {
			rabbit.BubbleState = THOUGHT_BUBBLE_BIG
		}
	} else if rabbit.BubbleState == THOUGHT_BUBBLE_FADE {
		if TimeSince(t, rabbit.Ticker) > THOUGHT_BUBBLE_DELAY {
			rabbit.BubbleState = THOUGHT_BUBBLE_SMALL
		}
	}
}

func (rabbit *Rabbit) PointInZone(x, y int32) bool {
	// calculate the distance
	dx := x - rabbit.MapX
	// add a little padding left of the rabbit too
	x_condition := dx >= -2 && dx <= THOUGHT_BUBBLE_ZONE
	y_condition := y-rabbit.MapY <= THOUGHT_BUBBLE_ZONE
	return x_condition && y_condition
}

func (rabbit *Rabbit) SetShowItem(t int32, status bool) {
	if !status {
		// need to reset the thought bubble animation after some delay
		rabbit.BubbleState = THOUGHT_BUBBLE_FADE
	}
	if rabbit.ShowItem != status {
		rabbit.Ticker = t
	}
	rabbit.ShowItem = status
}
