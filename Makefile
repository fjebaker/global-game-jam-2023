DOCKER_CMD := docker run --rm -v $(PWD):/src/cart -w /src/cart tinygo/tinygo:0.26.0
TINYGO := $(DOCKER_CMD) tinygo
GOFMT := $(DOCKER_CMD) gofmt
TIC80  := tic80

main:
	$(TINYGO) build -target ./target.json -o ./cart.wasm -panic print .

run: main
	$(TIC80) --skip --fs . --cmd 'load wasmdemo.wasmp & import binary cart.wasm & run & exit'

format:
	$(GOFMT) -w ./main.go ./tic80/tic80.go

cart: main
	rm -f game.tic
	$(TIC80) --skip --fs . --cmd 'load wasmdemo.wasmp & import binary cart.wasm & save game.tic & exit'

run-cart: cart
	$(TIC80) --fs . --cmd 'load game.tic & run & exit'

.PHONY: clean
clean:
	rm -f cart.wasm