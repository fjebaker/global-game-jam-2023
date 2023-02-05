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

const DESIREABLE_ITEM_NUMBER = 9

type RetrievableItem struct {
	Sprite        tic80.Sprite
	Bubble        ThoughtBubble
	ShowInTooltip bool
}

func NewRetrievableItem() RetrievableItem {
	return RetrievableItem{
		tic80.SquareSprite(int32(ITEM_RED_MUSHROOM), 1),
		NewThoughtBubble(),
		false,
	}
}

func (item *RetrievableItem) Id() int32 {
	return item.Sprite.Id
}

func (item *RetrievableItem) Update(t int32, player *Player, rabbit *Rabbit) {
	item.ShowInTooltip = player.HasItem
	is_in_zone := rabbit.PointInZone(player.X, player.Y)

	if player.HasItem && is_in_zone {
		player.HasItem = false
		item.Sprite.Id = int32(newDesiredItem())
		rabbit.HappySfx.Play()
		rabbit.ShowHeart = true
	}

	// propagate updates
	item.Bubble.Update(t, rabbit, is_in_zone)
}

func (item *RetrievableItem) Draw(t int32) {
	item.DrawTooltip()
	item.Bubble.Draw(t, &item.Sprite)
}

// utility

func newDesiredItem() DesireableItem {
	i := RandInt(0, DESIREABLE_ITEM_NUMBER)
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

func (item *RetrievableItem) DrawTooltip() {
	tic80.RectangleWithBorder(2, 2, 12, 12, 0, 9)
	if item.ShowInTooltip {
		item.Sprite.Draw(4, 4)
	}
}
