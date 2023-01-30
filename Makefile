DOCKER_CMD := docker run --rm -v $(PWD):/src/cart -w /src/cart tinygo/tinygo:0.26.0

TINYGO := $(DOCKER_CMD) tinygo
GOFMT := $(DOCKER_CMD) gofmt
TIC80  := tic80

WASMP_FILE := assets.wasmp

GOFLAGS := -target ./target.json -panic print -x  # -opt z -no-debug

CHECKFILESIZE = \
    FSIZE=$$(du -k cart.wasm | cut -f 1) ; \
    if [ $$FSIZE -gt 64 ]; then \
        >&2 echo "!!! filesize too big" ; exit 1 ; \
    fi

main:
	$(TINYGO) build $(GOFLAGS) -o ./cart.wasm .
	du -hs cart.wasm
	@$(CHECKFILESIZE)

run: main
	$(TIC80) --skip --fs . --cmd 'load $(WASMP_FILE) & import binary cart.wasm & run & exit'

.PHONY: format cart tic80 run-cart clean
format:
	$(GOFMT) -w ./main.go ./tic80/ ./cart/

cart: main
	rm -f game.tic
	$(TIC80) --skip --fs . --cmd 'load $(WASMP_FILE) & import binary cart.wasm & save game.tic & exit'

tic80:
	$(TIC80) --skip --fs . --cmd 'load $(WASMP_FILE) & import binary cart.wasm'

run-cart: cart
	$(TIC80) --fs . --cmd 'load game.tic & run & exit'

clean:
	rm -f cart.wasm