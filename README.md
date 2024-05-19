# go-spyder-game

This project implements a simple snake game (tested on linux-ubuntu-20.04LTS).

Some assumptions about this code:
- When the snake "eat" a fruit, it'll grow
- When the snake "eat" itself, it'll finish the game
- When the snake touch the boundaries, it'll be redirected to the oposite direction (like a portal)
- The food is randomly placed into the table
- When the user hits the `ESC` key, the game will stop.
- The speed (print delay) is fixed.
- The user can change the key once for each render view.

## How to run

```bash
# install packages
go mod download

# running as dev
go run main.go

# hwo to build the binary
go build -o snake main.go
```

![Captura de Tela (60)](https://github.com/ppcamp/go-snake-game/assets/38117637/ef2c79d1-6946-42c4-b84e-f7e64ef8d3eb)



**Why use external libraries?**

> To make it easier the acquisition of the keyboard (without needing to manage the handlers linux itself)
> To increasy the test readability basing on "clean code" minimalist styleguide


This repo use [git-commitzen](https://github.com/commitizen/cz-cli)
