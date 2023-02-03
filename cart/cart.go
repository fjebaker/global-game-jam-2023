package cart

import (
	"cart/tic80"
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
	tic80.Map(player.X, player.Y)

	player.HandleInteraction(t)
	player.Update(t)
	player.Draw(t)

	t = t + 1
	// avoid overflows
	t = t % 3600
}
