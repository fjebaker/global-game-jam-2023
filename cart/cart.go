package cart

import (
	"cart/tic80"
	"strconv"
)

var (
	mouse  tic80.MouseData
	t      int32
	player Player
)

func Start() {
	t = 1
	player = NewPlayer(20, 20)
}

// mainloop
func Loop() {
	tic80.Clear(0)
	mouse.Update()

	message := "Frame " + strconv.Itoa(int(t))
	tic80.Print(&message, 1, 1, 15, true, 1, false)

	player.HandleInteraction(t)
	player.Update(t)
	player.Draw(t)

	t = t + 1
	// avoid overflows
	t = t % 3600
}
