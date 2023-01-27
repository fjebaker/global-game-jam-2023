TINYGO := docker run --rm -v $(PWD):/src/cart -w /src/cart tinygo/tinygo:0.26.0 tinygo
TIC80  := tic80

main:
	$(TINYGO) build -target ./target.json -o ./cart.wasm -panic print .

run: main
	$(TIC80) --skip --fs . --cmd 'load wasmdemo.wasmp & import binary cart.wasm & run & exit'

cart: main
	rm -f game.tic
	$(TIC80) --skip --fs . --cmd 'load wasmdemo.wasmp & import binary cart.wasm & save game.tic & exit'
	$(TIC80) --fs . --cmd 'load game.tic & run & exit'

.PHONY: clean
clean:
	rm -f cart.wasm