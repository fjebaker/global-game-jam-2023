package cart

import "cart/tic80"

type DesireableItem int32

const (
	ITEM_GREEN_MUSHROOM DesireableItem = 32
	ITEM_RADISH         DesireableItem = 33
	ITEM_COIN           DesireableItem = 34
	ITEM_BLUE_MUSHROOM  DesireableItem = 35
	ITEM_RED_MUSHROOM   DesireableItem = 48
	ITEM_CLAY_POT       DesireableItem = 49
	ITEM_RUBY_RING      DesireableItem = 50
	ITEM_RED_THING      DesireableItem = 100
	ITEM_GREEN_THING    DesireableItem = 101
)

const (
	ITEM_FRAME_SIZE = 1
	ITEM_SCREEN_X   = 4
	ITEM_SCREEN_Y   = 4
)

const DESIREABLE_ITEM_COUNT = 9

type RetrievableItem struct {
	Sprite        tic80.Sprite
	Bubble        ThoughtBubble
	ShowInTooltip bool
}

///////////////////////////////////////////////////////////////////////////////
// Methods

func NewRetrievableItem() RetrievableItem {
	return RetrievableItem{
		tic80.SquareSprite(int32(ITEM_RED_MUSHROOM), ITEM_FRAME_SIZE),
		NewThoughtBubble(),
		false,
	}
}

func (item *RetrievableItem) Draw() {
	item.drawTooltip()
	item.Bubble.Draw(&item.Sprite)
}

func (item *RetrievableItem) Id() int32 {
	return item.Sprite.Id
}

func (item *RetrievableItem) Update(t int32, game *Game, player *Player, rabbit *Rabbit) {
	item.ShowInTooltip = player.HasItem
	is_in_zone := rabbit.PointInZone(player.X, player.Y)
	is_dead := rabbit.IsDead()

	if player.HasItem && is_in_zone && !is_dead {
		player.HasItem = false
		item.Sprite.Id = int32(newDesiredItem())
		rabbit.Heal(game)
	}

	// propagate updates
	item.Bubble.Update(t, rabbit, is_in_zone && !is_dead)
}

///////////////////////////////////////////////////////////////////////////////
// Utility

func (item *RetrievableItem) drawTooltip() {
	tic80.RectangleWithBorder(2, 2, 12, 12, 0, 9)
	if item.ShowInTooltip {
		item.Sprite.Draw(ITEM_SCREEN_X, ITEM_SCREEN_Y)
	}
}

func newDesiredItem() DesireableItem {
	i := RandInt(0, DESIREABLE_ITEM_COUNT)
	switch i {
	case 0:
		return ITEM_GREEN_MUSHROOM
	case 1:
		return ITEM_RADISH
	case 2:
		return ITEM_COIN
	case 3:
		return ITEM_BLUE_MUSHROOM
	case 4:
		return ITEM_RED_MUSHROOM
	case 5:
		return ITEM_CLAY_POT
	case 6:
		return ITEM_RUBY_RING
	case 7:
		return ITEM_RED_THING
	case 8:
		return ITEM_GREEN_THING
	// just incase
	default:
		return ITEM_RADISH
	}
}
