GO   = go
BIN  = isrss
ENTR = entr
SRC  = $(wildcard *.go)

$(BIN): $(SRC)
	$(GO) build .

run:
	$(GO) run .

watch:
	ls | $(ENTR) -r $(GO) run .

fmt:
	$(GO) fmt

clean:
	rm -f $(BIN)

.PHONY: run watch fmt clean
