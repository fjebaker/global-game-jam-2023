package tic80

import "unsafe"

type Sprite struct {
	Id               int32
	TransparentColor uint32
	Scale            int32
	Flip             TicFlip
	Rotate           TicRotate
	width            int32
	height           int32
}

func SquareSprite(id, width int32) Sprite {
	s := Sprite{
		Id:               id,
		width:            width,
		height:           width,
		TransparentColor: 0,
		Scale:            1,
	}
	return s
}

func (s *Sprite) Draw(x, y int32) {
	dx, dy := x, y
	if s.Scale > 1 {
		offset := s.width * 4 * s.Scale
		dx -= offset
		dy -= offset
	}

	_spr(
		s.Id,
		dx,
		dy,
		unsafe.Pointer(&s.TransparentColor),
		1,
		s.Scale,
		s.Flip,
		s.Rotate,
		s.width,
		s.height,
	)
}
