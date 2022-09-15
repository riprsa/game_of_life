package logic

import (
	"github.com/hararudoka/game_of_life/internal/config"
	"golang.org/x/exp/slices"
)

// Entry point of the game. Saves everything about logic here.
type Scene struct {
	// Map is a 2D array of cells that are true. All other cells are false so we don't need to store them.
	Map Map

	// LivingCells is map with all living cell. We use it fot logic and rendering.
	LivingCells map[Cord]struct{}

	// ToKeepAlive is a slice of the numbers of neighbors that a cell must have to stay alive.
	ToKeepAlive []int

	// ToBecomeAlive is a slice of the numbers of neighbors that a cell must have to become alive.
	ToBecomeAlive []int

	IsPaused bool

	// Speed controls number of loops in tick.
	Speed int

	tick int
}

// Creates new Scene from Config.
func NewScene(cfg config.Config) Scene {
	m := NewMap(cfg.MapSize, cfg.MapSize)

	return Scene{
		Map:           m,
		ToKeepAlive:   cfg.ToKeepAlive,
		ToBecomeAlive: cfg.ToBecomeAlive,
		LivingCells:   make(map[Cord]struct{}),
		IsPaused:      true,
		Speed:         5, // 1 - 5
		tick:          1,
	}
}

// UpdateScene changes the map for one tick.
func (s *Scene) UpdateScene() {
	if s.IsPaused {
		return
	}

	s.tick++
	if s.tick%(101-s.Speed*20) != 0 {
		return
	}

	newMap := NewMap(len(s.Map), len(s.Map[0]))

	newLivingCells := make(map[Cord]struct{})

	for cell := range s.LivingCells {
		if cell.X < 0 || cell.X >= len(s.Map) && cell.Y < 0 || cell.Y >= len(s.Map[0]) {
			continue
		}

		if slices.Contains(s.ToKeepAlive, len(s.Map.GetLivingNeighbors(cell.X, cell.Y))) {
			newMap[cell.X][cell.Y] = true
			newLivingCells[cell] = struct{}{}
		}

		neighbors := s.Map.GetDeadNeighbors(cell.X, cell.Y)
		for _, n := range neighbors {
			if slices.Contains(s.ToBecomeAlive, len(s.Map.GetLivingNeighbors(n.X, n.Y))) {
				if n.X < 0 || n.X >= len(newMap) || n.Y < 0 || n.Y >= len(newMap[0]) {
					continue
				}
				newMap[n.X][n.Y] = true
				newLivingCells[n] = struct{}{}
			}
		}
	}

	s.LivingCells = newLivingCells

	s.Map = newMap
}
