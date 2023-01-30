package main

import (
	"cart/cart"
	"cart/tic80"
)

//go:export BOOT
func BOOT() {
	tic80.Start()
	cart.Start()
}

//go:export TIC
func TIC() {
	cart.Loop()
}
