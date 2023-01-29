// some parts lifted from, others inspired by: https://github.com/sorucoder/tic80/blob/master/tic80.go
package tic80

import (
	"reflect"
	"unsafe"
)

// memory addresses
var (
	IO_RAM   = (*[0x18000]byte)(unsafe.Pointer(uintptr(0x00000)))
	FREE_RAM = (*[0x28000]byte)(unsafe.Pointer(uintptr(0x18000)))
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

//go:export cls
func cls(color int8)

//go:export print
func print(textBuffer unsafe.Pointer, x int32, y int32, color int8, fixed int8, scale int8, alt int8) int32

//go:linkname Init _start
func Init()

func Clear(color int8) {
	cls(color)
}

func Print(message *string, x int32, y int32, color int8, fixed int8, scale int8, alt int8) int32 {
	return print(toBytes(message), x, y, color, fixed, scale, alt)
}
