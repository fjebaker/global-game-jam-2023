package tic80

import (
	"unsafe"
)

func toBytes(s *string) unsafe.Pointer {
	// length 0, capacity len(*s) + 1
	data := make([]byte, 0, len(*s) + 2)
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

func Print(message *string, x, y, color int32, fixed bool, scale int32, alt bool) int32 {
	return _print(toBytes(message), x, y, color, fixed, scale, alt)
}

func (mouse *MouseData) Update() {
	_mouse(mouse)
}

func Trace(message *string, color int32) {
	_trace(toBytes(message), color)
}