package engine

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hararudoka/game_of_life/internal/config"
	"github.com/hararudoka/game_of_life/internal/logic"
	"golang.org/x/image/math/f64"

	_ "image/png"
)

type Game struct {
	Map          logic.Map
	world        *ebiten.Image
	Camera       Camera
	tilesImage   *ebiten.Image
	ScreenWidth  int
	ScreenHeight int
}

func NewGame(cfg config.Config) (*Game, error) {
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

	ebiten.SetMaxTPS(10) // ???
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")

	m := logic.NewMap(cfg.ToKeepAlive, cfg.ToBecomeAlive)

	return &Game{
		Camera:       Camera{ViewPort: f64.Vec2{float64(cfg.ScreenWidth), float64(cfg.ScreenHeight)}},
		Map:          m,
		world:        ebiten.NewImage(cfg.ScreenWidth, cfg.ScreenHeight),
		tilesImage:   ebiten.NewImageFromImage(img),
		ScreenWidth:  cfg.ScreenWidth,
		ScreenHeight: cfg.ScreenHeight,
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
	log.Println("Total Cells:", len(g.Map.Cells))

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x, y := float64(g.ScreenWidth/2), float64(g.ScreenHeight/2)

	g.world.Clear()

	for _, cell := range g.Map.Cells {
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(x+float64(cell.X), y+float64(cell.Y))

		g.world.DrawImage(g.tilesImage, op)
	}

	g.Camera.Render(g.world, screen)

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
