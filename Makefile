DOCKER_CMD := docker run --rm -v $(PWD):/src/cart -w /src/cart tinygo/tinygo:0.26.0

TINYGO := $(DOCKER_CMD) tinygo
GOFMT := $(DOCKER_CMD) gofmt
WASMOPT := $(DOCKER_CMD) wasm-opt
TIC80  := tic80

WASMP_FILE := assets.wasmp
HTML_ZIP := cart-html.zip
DEPLOY_BRANCH := github-pages
DEPLOY_FILES := index.html tic80.js tic80.wasm cart.tic

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
	$(TIC80) --cli --skip --fs . --cmd 'load $(WASMP_FILE) & import binary cart.wasm & save game.tic & exit'

tic80:
	$(TIC80) --skip --fs . --volume 6 --cmd 'load $(WASMP_FILE) & import binary cart.wasm'

run-cart: cart
	$(TIC80) --fs . --volume 6 --cmd 'load game.tic & run & exit'

publish: cart
	rm -f $(HTMLZIP)
	$(TIC80) --cli --fs . --skip --cmd 'load game.tic & export html $(HTML_ZIP) & exit'
	git switch $(DEPLOY_BRANCH)
	unzip -f $(HTML_ZIP)
	rm -f $(HTML_ZIP)
	git add $(DEPLOY_FILES)
	git commit -m "$$(git log --format=format:%H main -n 1)"
	git push origin $(DEPLOY_BRANCH)
	git switch main

clean:
	rm -f cart.wasm