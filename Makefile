DOCKER_CMD := docker run --rm -v $(PWD):/src/cart -w /src/cart tinygo/tinygo:0.26.0

TINYGO := $(DOCKER_CMD) tinygo
GOFMT := $(DOCKER_CMD) gofmt
WASMOPT := $(DOCKER_CMD) wasm-opt
TIC80  := tic80

WASMP_FILE := assets.wasmp

GOFLAGS := -target ./target.json -panic print -x -opt z -no-debug
WASMOPTFLAGS := --strip-debug --strip-dwarf -Oz --all-features

CHECKFILESIZE = \
    FSIZE=$$(du -k cart.wasm | cut -f 1) ; \
    if [ $$FSIZE -gt 64 ]; then \
        >&2 echo "!!! filesize too big" ; exit 1 ; \
    fi

main: format
	$(TINYGO) build $(GOFLAGS) -o ./cart.wasm .
	$(WASMOPT) $(WASMOPTFLAGS) -o ./cart.wasm ./cart.wasm
	du -hs cart.wasm
	@$(CHECKFILESIZE)

run: main
	$(TIC80) --skip --fs . --volume 6 --cmd 'load $(WASMP_FILE) & import binary cart.wasm & run & exit'

.PHONY: format cart tic80 run-cart clean
format:
	$(GOFMT) -w ./main.go ./tic80/ ./cart/

cart: main
	rm -f game.tic
	$(TIC80) --skip --fs . --volume 6 --cmd 'load $(WASMP_FILE) & import binary cart.wasm & save game.tic & exit'

tic80:
	$(TIC80) --skip --fs . --volume 6 --cmd 'load $(WASMP_FILE) & import binary cart.wasm'

run-cart: cart
	$(TIC80) --fs . --volume 6 --cmd 'load game.tic & run & exit'

clean:
	rm -f cart.wasm