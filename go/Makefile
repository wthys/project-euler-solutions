NAME:=euler
BIN_DIR:=./bin
PROG:=$(BIN_DIR)/$(NAME)

GOFILES:=$(shell find . -type f -name "*.go")

.PHONY: build clean run run-all

build: $(PROG)

$(PROG): go.mod $(GOFILES)
	go get
	go build -o $(PROG)

run: $(PROG)
	@$(PROG) run $(ARGS)

run-all: $(PROG)
	@for problem in `$(PROG) list`; do $(PROG) run $$problem; done

clean:
	rm go.sum
	rm -rf $(BIN_DIR)
