package engine

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hararudoka/game_of_life/internal/logic"
)

func (g *Game) gameplayControl() {
	// speed control
	if ebiten.IsKeyPressed(ebiten.Key1) {
		g.Scene.Speed = 1
	}
	if ebiten.IsKeyPressed(ebiten.Key2) {
		g.Scene.Speed = 2
	}
	if ebiten.IsKeyPressed(ebiten.Key3) {
		g.Scene.Speed = 3
	}
	if ebiten.IsKeyPressed(ebiten.Key4) {
		g.Scene.Speed = 4
	}
	if ebiten.IsKeyPressed(ebiten.Key5) {
		g.Scene.Speed = 5
	}

	// pause
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		if g.Scene.IsPaused {
			g.Scene.IsPaused = false
		} else {
			g.Scene.IsPaused = true
		}
	}

	// reset
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Scene.Map = logic.NewMap(len(g.Scene.Map), len(g.Scene.Map[0]))
	}
}

func (g *Game) cameraControl() {
	// movement
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.Camera.Position[0] -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.Camera.Position[0] += 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.Camera.Position[1] -= 3
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.Camera.Position[1] += 3
	}

	_, y := ebiten.Wheel()
	// zoom
	if y < 0 || ebiten.IsKeyPressed(ebiten.KeyQ) {
		if g.Camera.ZoomFactor > -15 {
			g.Camera.ZoomFactor += int(y - 1)
		}
	}

	if y > 0 || ebiten.IsKeyPressed(ebiten.KeyE) {
		if g.Camera.ZoomFactor < 300 {
			g.Camera.ZoomFactor += int(y + 1)
		}
	}
}

func (g *Game) cellEditor() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		worldX, worldY := g.Camera.ScreenToWorld(ebiten.CursorPosition())

		if int(worldX) < len(g.Scene.Map) && int(worldX) >= 0 && int(worldY) < len(g.Scene.Map[0]) && int(worldY) >= 0 {
			g.Scene.Map[int(worldX)][int(worldY)] = true
			// add to slice
			g.Scene.LivingCells[logic.Cord{X: int(worldX), Y: int(worldY)}] = struct{}{}
		}
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight) {
		worldX, worldY := g.Camera.ScreenToWorld(ebiten.CursorPosition())

		if int(worldX) < len(g.Scene.Map) && int(worldX) >= 0 && int(worldY) < len(g.Scene.Map[0]) && int(worldY) >= 0 {
			g.Scene.Map[int(worldX)][int(worldY)] = false
			delete(g.Scene.LivingCells, logic.Cord{X: int(worldX), Y: int(worldY)})
		}
	}
}
