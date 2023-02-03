package tic80

import "unsafe"

const (
	sHEIGHT int32 = 17
	sWIDTH        = 30
)
const (
	mHEIGHT int32 = 136
	mWIDTH  int32 = 240
)

func Map(x, y int32) {
	var safeX, safeY int32

	if x >= 0 && x <= (mWIDTH-sWIDTH) {
		safeX = x
	} else if x < 0 {
		safeX = 0
	} else {
		safeX = (mWIDTH - sWIDTH)
	}

	if y >= 0 && y <= (mHEIGHT-sHEIGHT) {
		safeY = y
	} else if y < 0 {
		safeY = 0
	} else {
		safeY = (mHEIGHT - sHEIGHT)
	}

	_map(
		safeX, safeY,
		// Full screen of tiles
		sWIDTH, sHEIGHT,
		// Alway from the top left corner
		0, 0,
		// Transparency
		unsafe.Pointer(nil), 0,
		// Scale
		1,
		// Unused
		0,
	)
}
