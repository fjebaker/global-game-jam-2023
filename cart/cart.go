package cart

import (
	"cart/tic80"
)

var (
	mouse  tic80.MouseData
	t      int32
	player Player
	world  World
)

func Start() {
	t = 1
	player = NewPlayer(95*8, PLAYER_OFFSET_Y+1)
	world = NewWorld(&player)
}

// mainloop
func Loop() {
	tic80.Clear(0)
	mouse.Update()

	player.HandleInteraction(t)
	player.Update(t, &world)
	world.Update(t, &player)

	world.Draw(t)
	player.Draw(t)

	t = t + 1
	// avoid overflows
	t = t % 3600
}
