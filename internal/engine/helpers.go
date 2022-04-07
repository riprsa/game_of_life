package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// load image from file
func loadImage(name string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile("assets/" + name)
	if err != nil {
		return nil, err
	}
	return img, nil
}
