package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"log"
	"math"

	"github.com/hararudoka/game_of_life/engine"

	_ "image/png"

	"golang.org/x/image/math/f64"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Camera struct {
	ViewPort   f64.Vec2
	Position   f64.Vec2
	ZoomFactor int
	Rotation   int
}

func (c *Camera) String() string {
	return fmt.Sprintf(
		"T: %.1f, R: %d, S: %d",
		c.Position, c.Rotation, c.ZoomFactor,
	)
}

func (c *Camera) viewportCenter() f64.Vec2 {
	return f64.Vec2{
		c.ViewPort[0] * 0.5,
		c.ViewPort[1] * 0.5,
	}
}

func (c *Camera) worldMatrix() ebiten.GeoM {
	m := ebiten.GeoM{}
	m.Translate(-c.Position[0], -c.Position[1])
	// We want to scale and rotate around center of image / screen
	m.Translate(-c.viewportCenter()[0], -c.viewportCenter()[1])
	m.Scale(
		math.Pow(1.01, float64(c.ZoomFactor)),
		math.Pow(1.01, float64(c.ZoomFactor)),
	)
	m.Rotate(float64(c.Rotation) * 2 * math.Pi / 360)
	m.Translate(c.viewportCenter()[0], c.viewportCenter()[1])
	return m
}

func (c *Camera) Render(world, screen *ebiten.Image) {
	screen.DrawImage(world, &ebiten.DrawImageOptions{
		GeoM: c.worldMatrix(),
	})
}

func (c *Camera) ScreenToWorld(posX, posY int) (float64, float64) {
	inverseMatrix := c.worldMatrix()
	if inverseMatrix.IsInvertible() {
		inverseMatrix.Invert()
		return inverseMatrix.Apply(float64(posX), float64(posY))
	} else {
		// When scaling it can happend that matrix is not invertable
		return math.NaN(), math.NaN()
	}
}

func (c *Camera) Reset() {
	c.Position[0] = 0
	c.Position[1] = 0
	c.Rotation = 0
	c.ZoomFactor = 0
}

type Game struct {
	Map          engine.Map
	world        *ebiten.Image
	camera       Camera
	tilesImage   *ebiten.Image
	screenWidth  int
	screenHeight int
}

func NewGame(worldWidth, worldHeight int) *Game {
	b64 := "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAQAAAC1HAwCAAAAC0lEQVR42mP8/x8AAwMCAO+ip1sAAAAASUVORK5CYII="

	// base64 to image
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		log.Fatal(err)
	}

	return &Game{
		camera:       Camera{ViewPort: f64.Vec2{float64(worldWidth), float64(worldHeight)}, ZoomFactor: 300},
		Map:          engine.NewMap(),
		world:        ebiten.NewImage(worldWidth, worldHeight),
		tilesImage:   ebiten.NewImageFromImage(img),
		screenWidth:  worldWidth,
		screenHeight: worldHeight,
	}
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.camera.Position[0] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.camera.Position[0] += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.camera.Position[1] -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.camera.Position[1] += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		g.camera.ZoomFactor -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) {
		if g.camera.ZoomFactor < 2400 {
			g.camera.ZoomFactor += 1
		}
	}

	g.Map.UpdateMap()

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	x, y := float64(g.screenWidth/2), float64(g.screenHeight/2)

	g.world.Clear()

	for _, cell := range g.Map {
		op := &ebiten.DrawImageOptions{}

		op.GeoM.Translate(x+float64(cell.X), y+float64(cell.Y))

		g.world.DrawImage(g.tilesImage, op)
	}

	g.camera.Render(g.world, screen)

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.screenWidth, g.screenHeight
}

func main() {
	g := NewGame(800, 800)

	ebiten.SetMaxTPS(10)

	ebiten.SetWindowSize(g.screenWidth, g.screenHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
