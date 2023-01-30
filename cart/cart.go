package cart

import (
	"cart/tic80"
	"strconv"
	"math/rand"
)

var (
	mouse           tic80.MouseData
	t               int
	x, y            int32
	moving          bool
	player, monster tic80.Sprite
	count, direction int32
	m_x, m_y		int32
)

func randInt(min, max int) int {
    return min + rand.Intn(max-min)
}

func Start() {
	t = 1
	count = 0
	direction = 0
	x = 50
	y = 50
	moving = false
	player = tic80.SquareSprite(258, 1)
	monster = tic80.SquareSprite(272, 2)
	m_x = 100
	m_y = 100
	// tic80.Music(0, 0, 0, true, false, 100, 6)
}

const IDLE1 = 257
const IDLE2 = 258
const LEFT1 = 259
const LEFT2 = 260

const (
	M_IDLE1 int32 = 272 + (iota * 2)
	M_IDLE2
	M_LEFT1
	M_LEFT2
	M_LEFT3
	M_LEFT4
)

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

func getMonsterIdle(t int) int32 {
	if t%60 > 30 {
		return M_IDLE1
	} else {
		return M_IDLE2
	}
}

func getMonsterWalk(t int) int32 {
	i := t % 120
	switch {
	case i <= 30:
		return M_LEFT1
	case i <= 60:
		return M_LEFT2
	case i <= 90:
		return M_LEFT3
	default:
		return M_LEFT4
	}
}

// mainloop
func Loop() {
	tic80.Clear(0)
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

	monster.Id = getMonsterIdle(t)

	monster.Draw(100, 50)
	player.Draw(x, y)


	moving = false
	t = t + 1
	// avoid overflows
	t = t % 3600
}
