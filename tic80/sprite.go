package tic80

import "unsafe"

type Sprite struct {
	Id                    int32
	TransparentColor      uint32
	TransparentColorCount uint32
	Scale                 int32
	Flip                  TicFlip
	Rotate                TicRotate
	width                 int32
	height                int32
}

func SquareSprite(id, width int32) Sprite {
	s := Sprite{
		Id:                    id,
		width:                 width,
		height:                width,
		TransparentColor:      0,
		TransparentColorCount: 1,
		Scale:                 1,
	}
	return s
}

func (s *Sprite) Draw(x, y int32) {
	_spr(
		s.Id,
		x,
		y,
		unsafe.Pointer(&s.TransparentColor),
		s.TransparentColorCount,
		s.Scale,
		s.Flip,
		s.Rotate,
		s.width,
		s.height,
	)
}
