package engine

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hararudoka/game_of_life/internal/logic"
	"golang.org/x/image/math/f64"
)

type Game struct {
	Map          logic.Map
	World        *ebiten.Image
	Camera       Camera
	tilesImage   *ebiten.Image
	ScreenWidth  int
	ScreenHeight int
}

func NewGame(worldWidth, worldHeight int) (*Game, error) {
	b64 := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO+ip1sAAAAASUVORK5CYII="

	// base64 to image
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	return &Game{
		Camera:       Camera{ViewPort: f64.Vec2{float64(worldWidth), float64(worldHeight)}, ZoomFactor: 300},
		Map:          logic.NewMap(),
		World:        ebiten.NewImage(worldWidth, worldHeight),
		tilesImage:   ebiten.NewImageFromImage(img),
		ScreenWidth:  worldWidth,
		ScreenHeight: worldHeight,
	}, nil
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.Camera.Position[0] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Camera.Position[0] += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.Camera.Position[1] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.Camera.Position[1] += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.Camera.ZoomFactor -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		if g.Camera.ZoomFactor < 2400 {
			g.Camera.ZoomFactor += 1
		}
	}

	g.Map.UpdateMap()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x, y := float64(g.ScreenWidth/2), float64(g.ScreenHeight/2)

	g.World.Clear()

	for _, cell := range g.Map {
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(x+float64(cell.X), y+float64(cell.Y))

		g.World.DrawImage(g.tilesImage, op)
	}

	g.Camera.Render(g.World, screen)

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
