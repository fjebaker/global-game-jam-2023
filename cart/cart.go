package cart

import (
	"cart/tic80"
	"strconv"
)

var (
	mouse  tic80.MouseData
	t      int
	x, y   int32
	moving bool
	player tic80.Sprite
)

func Start() {
	t = 1
	x = 50
	y = 50
	moving = false
	player = tic80.SquareSprite(258, 1)
}

const IDLE1 = 257
const IDLE2 = 258
const LEFT1 = 259
const LEFT2 = 260

func getMotion(t int) int32 {
	if t%30 > 15 {
		return LEFT1
	} else {
		return LEFT2
	}
}

func getIdle(t int) int32 {
	if t%30 > 15 {
		return IDLE1
	} else {
		return IDLE2
	}
}

// mainloop
func Loop() {
	tic80.Clear(13)
	mouse.Update()

	message := "Frame " + strconv.Itoa(t)
	tic80.Print(&message, 1, 1, 15, true, 1, false)

	if tic80.BUTTON_UP.IsPressed() {
		y = y - 1
		player.Id = getMotion(t)
		moving = true
	}
	if tic80.BUTTON_DOWN.IsPressed() {
		y = y + 1
		player.Id = getMotion(t)
		moving = true
	}
	if tic80.BUTTON_LEFT.IsPressed() {
		x = x - 1
		moving = true
		player.Id = getMotion(t)
		player.Flip = 0
	}
	if tic80.BUTTON_RIGHT.IsPressed() {
		x = x + 1
		moving = true
		player.Id = getMotion(t)
		player.Flip = 1
	}
	if moving == false {
		player.Id = getIdle(t)
	}

	player.Draw(x, y)

	moving = false
	t = t + 1
	// avoid overflows
	t = t % 3600
}
