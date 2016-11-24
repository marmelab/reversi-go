.PHONY: install run debug test lint benchmark

GO_BIN := docker run \
	-p 8080:8080 \
	--interactive \
	--rm \
	--tty \
	--volume="$(CURDIR):/srv" \
	reversi-go

# Initialization ################

install:
	docker build --tag=reversi-go .
	$(GO_BIN) go get github.com/fatih/color
	$(GO_BIN) go get github.com/olekukonko/tablewriter
	$(GO_BIN) go get github.com/hishboy/gocommons/lang

# Run ###########################

run:
	$(GO_BIN) go run src/reversi/main.go

# Run Server ####################

run-server:
	$(GO_BIN) go run src/reversi/ai/server/server.go

# Run With debug trace ##########

debug:
	$(GO_BIN) go run src/reversi/main.go --debugfile=debug.log

# Tests #########################

test:
	$(GO_BIN) bash -c "cd src/reversi && go test ./..."

# Lint ##########################

lint:
	$(GO_BIN) gofmt -w src/reversi

# Bench #########################

benchmark:
	$(GO_BIN) bash -c "cd src/reversi/ai && go test -run=XXX -bench=."
