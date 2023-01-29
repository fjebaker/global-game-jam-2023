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

// still need this since _start calls main
func main() {}
