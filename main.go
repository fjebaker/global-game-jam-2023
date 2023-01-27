// inspired by https://github.com/sorucoder/tic80/blob/master/tic80.go
package main

import "cart/tic80"

//go:export TIC
func TIC() {
	tic80.Clear(13)
	message := "Hello World"
    tic80.Print(&message, 60, 84, 15, 1, 1, 0)
}

//go:export BOOT
func BOOT() {
	tic80.Init()
}

// still need this for some odd reason
func main() {}