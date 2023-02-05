package cart

import "cart/tic80"

type ThoughtBubbleState int32

const (
	THOUGHT_BUBBLE_FADE ThoughtBubbleState = iota
	THOUGHT_BUBBLE_SMALL
	THOUGHT_BUBBLE_BIG
)

const ITEM_OFFSET = 4

const (
	THOUGHT_BUBBLE_DELAY = 60
)

type ThoughtBubble struct {
	State     ThoughtBubbleState
	StartTime int32
	X, Y      int32
	Show      bool
}

func NewThoughtBubble() ThoughtBubble {
	return ThoughtBubble{THOUGHT_BUBBLE_BIG, 0, 0, 0, false}
}

func (bubble *ThoughtBubble) Update(t int32, rabbit *Rabbit, show bool) {
	if bubble.Show != show {
		bubble.StartTime = t
	}
	bubble.Show = show

	bubble.X = rabbit.X - ITEM_OFFSET
	bubble.Y = rabbit.Y - ITEM_OFFSET

	// dT := TimeSince(t, bubble.StartTime)

	// is_time := dT > THOUGHT_BUBBLE_DELAY
	// is_growing := bubble.State == THOUGHT_BUBBLE_SMALL && is_time
	// is_fading := bubble.State == THOUGHT_BUBBLE_FADE && is_time

	// switch {
	// case bubble.Show && bubble.State == THOUGHT_BUBBLE_FADE:
	// 	bubble.State = THOUGHT_BUBBLE_SMALL
	// case !bubble.Show && is_fading:
	// 	// go from big to small
	// 	bubble.State = THOUGHT_BUBBLE_SMALL
	// 	// we still show
	// 	bubble.Show = true
	// case bubble.Show && (is_growing || is_fading):
	// 	// if we are fading and player comes back, go back to displaying big
	// 	bubble.State = THOUGHT_BUBBLE_BIG
	// }
}

func (bubble *ThoughtBubble) Draw(t int32, item *tic80.Sprite) {
	x_item := bubble.X
	y_item := bubble.Y

	if bubble.Show {
		if bubble.State >= THOUGHT_BUBBLE_SMALL {
			tic80.EllipseWithBorder(x_item+10, y_item+11, 2, 1, 12, 13)
		}
		if bubble.State >= THOUGHT_BUBBLE_BIG {
			tic80.EllipseWithBorder(x_item+3, y_item+3, 7, 6, 12, 13)
			item.Draw(x_item, y_item)
		}
	}
}
