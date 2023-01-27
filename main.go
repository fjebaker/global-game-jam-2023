package main

// #include <stdlib.h>
import "C"
import (
	"unsafe"
)

//go:export cls
func cls(color int32)

//go:export print
func print(txt *C.char, x int32, y int32, color int32, fixed int32, scale int32, alt int32) int32

//go:export TIC
func TIC() {
	cls(13)
	message := C.CString("Hello World")
	defer C.free(unsafe.Pointer(message))
	print(message, 3, 3, 15, 0, 1, 1);
}