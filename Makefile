.PHONY: install run test lint benchmark

GO_BIN := docker run \
	--interactive \
	--rm \
	--tty \
	-P \
	--volume="$(CURDIR):/srv" \
	reversi-go

# Initialization ################

install:
	docker build --tag=reversi-go .
	$(GO_BIN) go get github.com/fatih/color
	$(GO_BIN) go get github.com/olekukonko/tablewriter

# Run ###########################

run:
	$(GO_BIN) go run src/reversi/main.go

# Tests #########################

test:
	$(GO_BIN) bash -c "cd src/reversi/game/player && go test"
	$(GO_BIN) bash -c "cd src/reversi/game/vector && go test"
	$(GO_BIN) bash -c "cd src/reversi/game/cell && go test"
	$(GO_BIN) bash -c "cd src/reversi/game/board && go test"
	$(GO_BIN) bash -c "cd src/reversi/game/game && go test"
	$(GO_BIN) bash -c "cd src/reversi/ai && go test"
	$(GO_BIN) bash -c "cd src/reversi/ai/scoring && go test"

# Lint ##########################

lint:
	$(GO_BIN) gofmt -w src/

# Bench #########################

benchmark:
	$(GO_BIN) bash -c "cd src/reversi/ai && go test -run=XXX -bench=."
