TINYGO := docker run --rm -v $(PWD):/home/tinygo/src -w /home/tinygo/src tinygo/tinygo:0.26.0 tinygo

main:
	$(TINYGO) build -target ./target.json -o ./cart.wasm .

run: main
	tic80 --skip --fs . --cmd 'load wasmdemo.wasmp & import binary cart.wasm & run & exit'

cart: main
	rm -f game.tic
	tic80 --skip --fs . --cmd 'load wasmdemo.wasmp & import binary cart.wasm & save game.tic & exit'
	tic80 --fs . --cmd 'load game.tic & run & exit'

.PHONY: clean
clean:
	rm -f cart.wasm