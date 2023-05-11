package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"rope-simulation/rope"
)

func main() {
	game, err := rope.NewGame()

	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(rope.ScreenWidth, rope.ScreenHeight)
	ebiten.SetWindowTitle("Rope simulation (demo)")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
