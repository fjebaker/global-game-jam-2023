package main

import (
	"cart/cart"
)

//go:export BOOT
func BOOT() {
	cart.Start()
}

//go:export TIC
func TIC() {
	cart.Loop()
}
