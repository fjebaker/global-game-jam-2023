package cart

import (
	"cart/tic80"
	// "strconv"
)

var (
	mouse tic80.MouseData
	t int
)

func Start() {
	tic80.Start()
	t = 1
}

// mainloop
func Loop() {
	tic80.Clear(13)
	mouse.Update()
	// message := strconv.Itoa(t)
	// _ = message
	// tic80.Print(&message, 60, 84, 15, 1, 1, 0)
	// tic80.Trace(&message, 1)

	t = t + 1
	// avoid overflows
	t = t % 60
}
