// some functions lifted from, others inspired by: https://github.com/sorucoder/tic80/blob/master/tic80.go
package tic80

import (
	"unsafe"
)

// memory addresses
var (
	IO_RAM   = (*[0x18000]byte)(unsafe.Pointer(uintptr(0x00000)))
	FREE_RAM = (*[0x28000]byte)(unsafe.Pointer(uintptr(0x18000)))
)

// ButtonCode represents a button id
type ButtonCode int32

// Gamepads
const (
	GAMEPAD_1 ButtonCode = 8 * iota
	GAMEPAD_2
	GAMEPAD_3
	GAMEPAD_4
)

// Buttons
const (
	BUTTON_UP ButtonCode = iota
	BUTTON_DOWN
	BUTTON_LEFT
	BUTTON_RIGHT
	BUTTON_A
	BUTTON_B
	BUTTON_X
	BUTTON_Y
)

// KeyCode represents a keyboard id
type KeyCode int

// Keyboard keys.
const (
	KEY_A KeyCode = iota + 1
	KEY_B
	KEY_C
	KEY_D
	KEY_E
	KEY_F
	KEY_G
	KEY_H
	KEY_I
	KEY_J
	KEY_K
	KEY_L
	KEY_M
	KEY_N
	KEY_O
	KEY_P
	KEY_Q
	KEY_R
	KEY_S
	KEY_T
	KEY_U
	KEY_V
	KEY_W
	KEY_X
	KEY_Y
	KEY_Z
	KEY_ZERO
	KEY_ONE
	KEY_TWO
	KEY_THREE
	KEY_FOUR
	KEY_FIVE
	KEY_SIX
	KEY_SEVEN
	KEY_EIGHT
	KEY_NINE
	KEY_MINUS
	KEY_EQUALS
	KEY_LEFTBRACKET
	KEY_RIGHTBRACKET
	KEY_BACKSLASH
	KEY_SEMICOLON
	KEY_APOSTROPHE
	KEY_GRAVE
	KEY_COMMA
	KEY_PERIOD
	KEY_SLASH
	KEY_SPACE
	KEY_TAB
	KEY_RETURN
	KEY_BACKSPACE
	KEY_DELETE
	KEY_INSERT
	KEY_PAGEUP
	KEY_PAGEDOWN
	KEY_HOME
	KEY_END
	KEY_UP
	KEY_DOWN
	KEY_LEFT
	KEY_RIGHT
	KEY_CAPSLOCK
	KEY_CTRL
	KEY_SHIFT
	KEY_ALT
)

//go:export btn
func _btn(id int32) int32

//go:export btnp
func _btnp(id, hold, period int32) bool

//go:export clip
func _clip(x, y, width, height int32)

//go:export cls
func _cls(color int32)

//go:export circ
func _circ(x, y, radius int32, color int32)

//go:export circb
func _circb(x, y, radius int32, color int32)

//go:export elli
func _elli(x, y, radiusX, radiusY int32, color int32)

//go:export ellib
func _ellib(x, y, radiusX, radiusY int32, color int32)

//go:export exit
func _exit()

//go:export fget
func _fget(sprite int32, flag int8) bool

//go:export fset
func _fset(sprite int32, flag int8, value bool)

//go:export font
func _font(textBuffer unsafe.Pointer, x, y int32, transparentColorBuffer unsafe.Pointer, transparentColorCount int8, characterWidth, characterHeight int8, fixed bool, scale int8, useAlternateFontPage bool) int32

//go:export key
func _key(id int32) int32

//go:export keyp
func _keyp(id int8, hold, period int32) int32

//go:export line
func _line(x0, y0, x1, y1 float32, color int32)

//go:export map
func _map(x, y, width, height, screenX, screenY int32, transparentColorBuffer unsafe.Pointer, transparentColorCount int8, unused int32)

//go:export memcpy
func _memcpy(destination, source, length int32)

//go:export memset
func _memset(address, value, length int32)

//go:export mget
func _mget(x, y int32) int32

//go:export mset
func _mset(x, y, value int32)

type MouseData struct {
	x       int16
	y       int16
	scrollX int8
	scrollY int8
	left    bool
	middle  bool
	right   bool
}

//go:export mouse
func _mouse(data *MouseData)

//go:export music
func _music(track_id, frame, row int32, loop, sustain bool, tempo, speed int32)

//go:export print
func _print(text unsafe.Pointer, x, y, color int32, fixed bool, scale int32, alt bool) int32

//go:export peek
func _peek(address int32, bits int8) int8

//go:export pix
func _pix(x, y int32, color int32) uint8

//go:export pmem
func _pmem(address int32, value int64) uint32

//go:export poke
func _poke(address int32, value, bits int8)

//go:export rect
func _rect(x, y, width, height int32, color int32)

//go:export rectb
func _rectb(x, y, width, height int32, color int32)

//go:export sfx
func _sfx(id, note, octave, duration, channel, volumeLeft, volumeRight, speed int32)

//go:export spr
func _spr(id, x, y int32, transparentColorBuffer unsafe.Pointer, transparentColorCount uint32, scale, flip, rotate, width, height int32)

//go:export sync
func _sync(mask int32, bank, toCart int8)

//go:export ttri
func _ttri(x0, y0, x1, y1, x2, y2, u0, v0, u1, v1, u2, v2 float32, useTiles int32, transparentColorBuffer unsafe.Pointer, transparentColorCount int8, z0, z1, z2 float32, depth bool)

//go:export time
func _time() float32

//go:export trace
func _trace(messageBuffer unsafe.Pointer, color int32)

//go:export tri
func _tri(x0, y0, x1, y1, x2, y2 float32, color int32)

//go:export trib
func _trib(x0, y0, x1, y1, x2, y2 float32, color int32)

//go:export tstamp
func _tstamp() uint32

//go:linkname Start _start
func Start()

//go:export main.main
func main() {}
