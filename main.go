package main

import (
	"log"

	"github.com/hararudoka/game_of_life/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := engine.NewGame(800, 800)

	ebiten.SetMaxTPS(10)

	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
