# reversi-go

Turn based reversi game for command-line in Golang

## Install

Install docker image

```
make install
```

## Run

Run the game

```
make run
```

## Run Server

Run the game AI server
The game is exposed on the 8080 port. To use it, you should make a GET request with ?type=<your-type> (1 (black) / 2 (white)) and the board in JSON (eg: [[0, 0, 0, 0], [0, 1, 1, ...], ...]) as body.

```
make run-server
```

## Test

Runs all the tests of the project

```
make test
```

## Lint

Runs code linter on the project

```
make lint
```
