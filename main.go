package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/wonderbird/golabyrinth/labyrinth"
)

func main() {
	game := &labyrinth.Game{}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
