package tic80

import (
	"reflect"
	"unsafe"
)

func toByteData(goBytes *[]byte) (buffer unsafe.Pointer, count int) {
	if goBytes != nil {
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(goBytes))
		buffer = unsafe.Pointer(sliceHeader.Data)
		// See https://tinygo.org/docs/guides/compatibility/#reflectsliceheader-and-reflectstringheader.
		count = int(sliceHeader.Len)
	}
	return
}

func toBytes(goString *string) unsafe.Pointer {
	data := new([]byte)
	*data = make([]byte, 0, len(*goString)+1)
	for _, token := range *goString {
		if token > 0 {
			switch {
			case token <= 0x7F:
				*data = append(*data, byte(token))
			default:
				*data = append(*data, byte('?'))
			}
		}
	}
	*data = append(*data, 0)
	buffer, _ := toByteData(data)
	return buffer
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