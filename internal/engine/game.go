package engine

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hararudoka/game_of_life/internal/config"
	logic "github.com/hararudoka/game_of_life/internal/logic"
	"golang.org/x/image/math/f64"

	_ "image/png"
)

type Game struct {
	Scene  logic.Scene
	world  *ebiten.Image
	Camera Camera

	liveCellsImage *ebiten.Image
	deadCellsImage *ebiten.Image

	ScreenWidth  int
	ScreenHeight int
}

func NewGame(cfg config.Config) (*Game, error) {
	liveCell, err := loadImage("black_1x1.png")
	if err != nil {
		return nil, err
	}

	deadCell, err := loadImage("white_1x1.png")
	if err != nil {
		return nil, err
	}

	ebiten.SetMaxTPS(cfg.TPS)
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")

	m := logic.NewScene(cfg)

	// size of world
	w := ebiten.NewImage(cfg.MapSize, cfg.MapSize)
	w.Fill(color.White)

	return &Game{
		Camera:         Camera{ViewPort: f64.Vec2{float64(cfg.ScreenWidth / 2), float64(cfg.ScreenHeight / 2)}, ZoomFactor: 150},
		Scene:          m,
		world:          w,
		liveCellsImage: liveCell,
		deadCellsImage: deadCell,
		ScreenWidth:    cfg.ScreenWidth,
		ScreenHeight:   cfg.ScreenHeight,
	}, nil
}

func (g *Game) Update() error {
	// cell edition functionality
	g.cellEditor()

	// camera movement and zoom
	g.cameraControl()

	// common game control
	g.gameplayControl()

	// update the scene
	g.Scene.UpdateScene()

	return nil
}

// render the scene. very fast
func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Fill(color.White)

	// draw living cells to map.
	for cell := range g.Scene.LivingCells {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(cell.X), float64(cell.Y))
		g.world.DrawImage(g.liveCellsImage, op)
	}

	g.Camera.Render(g.world, screen)

	// GUI (debug)
	worldX, worldY := g.Camera.ScreenToWorld(ebiten.CursorPosition())
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("x: %.2f, y: %.2f", worldX, worldY),
		0, g.ScreenHeight-16,
	)
	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f\nSpeed: %v\nLiving Cells: %v\n", ebiten.CurrentTPS(), g.Scene.Speed, len(g.Scene.LivingCells)),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.ScreenWidth, g.ScreenHeight
}
