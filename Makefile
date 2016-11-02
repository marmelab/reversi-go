.PHONY: install run test lint

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

# Run ###########################

run:
	$(GO_BIN) go run src/reversi/main.go

# Tests #########################

test:
	$(GO_BIN) bash -c "cd src/reversi/game/cell && go test"
	$(GO_BIN) bash -c "cd src/reversi/game/board && go test"
	$(GO_BIN) bash -c "cd src/reversi/game/player && go test"
	$(GO_BIN) bash -c "cd src/reversi/game/game && go test"

# Lint ##########################

lint:
	$(GO_BIN) go get github.com/golang/lint/golint
	$(GO_BIN) ./bin/golint src/
