package main

import (
	"log"

	"github.com/hararudoka/game_of_life/internal/config"
	"github.com/hararudoka/game_of_life/internal/engine"

	"github.com/hajimehoshi/ebiten/v2"
)

// TODO:
// - game interface with comfort buttons
// - features like play, pause, speed up, slow down, etc.
// - fix lags with more than 2K-6K cells

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	g, err := engine.NewGame(cfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
