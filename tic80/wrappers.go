package tic80

import (
	"unsafe"
)

func toBytes(s *string) unsafe.Pointer {
	// length 0, capacity len(*s) + 1
	data := make([]byte, 0, len(*s)+2)
	for _, token := range *s {
		switch {
		case token <= 0:
			break
		case token <= 0x7F:
			data = append(data, byte(token))
		default:
			data = append(data, byte('?'))
		}
	}
	data = append(data, 0)
	return unsafe.Pointer(&data[0])
}

func Clear(color int32) {
	_cls(color)
}

func FGet(slot int32, flagIndex int8) bool {
	return _fget(slot, flagIndex)
}

func Music(track_id, frame, row int32, loop, sustain bool, tempo, speed int32) {
	_music(track_id, frame, row, loop, sustain, tempo, speed)
}

func Print(message *string, x, y, color int32, fixed bool, scale int32, alt bool) int32 {
	return _print(toBytes(message), x, y, color, fixed, scale, alt)
}

func Trace(message *string, color int32) {
	_trace(toBytes(message), color)
}

func (mouse *MouseData) Update() {
	_mouse(mouse)
}

func (id ButtonCode) IsPressed() bool {
	return _btn(int32(id)) > 0
}

func PaintPixel(x, y, color int32) {
	_pix(x, y, color)
}

func Rectangle(x, y, width, height, color int32) {
	_rect(x, y, width, height, color)
}

func RectangleBorder(x, y, width, height, color int32) {
	_rectb(x, y, width, height, color)
}

func RectangleWithBorder(x, y, width, height, color, border_color int32) {
	Rectangle(x, y, width, height, color)
	RectangleBorder(x-1, y-1, width+2, height+2, border_color)
}

func Ellipse(x, y, radius_x, radius_y, color int32) {
	_elli(x, y, radius_x, radius_y, color)
}

func EllipseBorder(x, y, radius_x, radius_y, color int32) {
	_ellib(x, y, radius_x, radius_y, color)
}

func EllipseWithBorder(x, y, radius_x, radius_y, color, border_color int32) {
	Ellipse(x, y, radius_x, radius_y, color)
	EllipseBorder(x, y, radius_x, radius_y, border_color)
}
