package logic

import (
	"golang.org/x/exp/slices"
)

func init() {
}

// Cell is a struct that contains the x and y cordinates of a cell.
type Cell struct {
	X int
	Y int
}

type Cells []Cell

type Map struct {
	// Cells is a 2D array of cells that are true. All other cells are false so we don't need to store them.
	Cells Cells

	// ToKeepAlive is a slice of ints that are the number of neighbors that a cell must have to stay alive.
	ToKeepAlive []int

	// ToBecomeAlive is a slice of ints that are the number of neighbors that a cell must have to become alive.
	ToBecomeAlive []int
}

// NewMap returns a new map of cells.
func NewMap(toKeepAlive, toBecomeAlive []int) Map {
	c := make(Cells, 0)

	// rn i dont have time to rewrite this to normal input or ingame clicks, so i will just hardcode it.
	c = append(c, Cell{X: 0, Y: 0})
	c = append(c, Cell{X: 1, Y: 0})
	c = append(c, Cell{X: 1, Y: 1})
	c = append(c, Cell{X: 2, Y: 1})
	c = append(c, Cell{X: 0, Y: 2})

	return Map{
		Cells:         c,
		ToKeepAlive:   toKeepAlive,
		ToBecomeAlive: toBecomeAlive,
	}
}

// GetLivingNeighbors returns a slice of cells that are adjacent to the cell passed in.
func (cs Cells) GetLivingNeighbors(c Cell) []Cell {
	var neighbors []Cell
	for _, cell := range cs {
		switch {
		case cell.X == c.X+1 && cell.Y == c.Y:
			neighbors = append(neighbors, cell)
		case cell.X == c.X-1 && cell.Y == c.Y:
			neighbors = append(neighbors, cell)
		case cell.X == c.X && cell.Y == c.Y+1:
			neighbors = append(neighbors, cell)
		case cell.X == c.X && cell.Y == c.Y-1:
			neighbors = append(neighbors, cell)
		case cell.X == c.X+1 && cell.Y == c.Y+1:
			neighbors = append(neighbors, cell)
		case cell.X == c.X-1 && cell.Y == c.Y-1:
			neighbors = append(neighbors, cell)
		case cell.X == c.X+1 && cell.Y == c.Y-1:
			neighbors = append(neighbors, cell)
		case cell.X == c.X-1 && cell.Y == c.Y+1:
			neighbors = append(neighbors, cell)
		}
	}
	return neighbors
}

// GetDeadNeighbors returns a slice of cells that are adjacent to the cell passed in and not in Map.
func (cs Cells) GetDeadNeighbors(c Cell) []Cell {
	// All neighbors of the cell.
	neighbors := []Cell{
		{X: c.X + 1, Y: c.Y},
		{X: c.X - 1, Y: c.Y},
		{X: c.X, Y: c.Y + 1},
		{X: c.X, Y: c.Y - 1},
		{X: c.X + 1, Y: c.Y + 1},
		{X: c.X - 1, Y: c.Y - 1},
		{X: c.X + 1, Y: c.Y - 1},
		{X: c.X - 1, Y: c.Y + 1},
	}

	// сandidates are the neighbors that are dead now.
	var сandidates []Cell

	for _, neighbor := range neighbors {
		if slices.Index(cs, neighbor) == -1 {
			сandidates = append(сandidates, neighbor)
		}
	}
	return сandidates
}

// RemoveDubles removes duplicates from a Map.
func (cs *Cells) RemoveDuplicates() {
	var uniqueMap Cells
	u := make(map[Cell]struct{})
	for _, c := range *cs {
		u[c] = struct{}{}
	}
	for k := range u {
		uniqueMap = append(uniqueMap, k)
	}
	*cs = uniqueMap
}

// DeleteCell removes a cell from the map.
// func (cs *Cells) DeleteCell(c Cell) {
// 	for i, cell := range *cs {
// 		if cell.X == c.X && cell.Y == c.Y {
// 			*cs = append((*cs)[:i], (*cs)[i+1:]...)
// 			return
// 		}
// 	}
// }

// UpdateMap changes the map for one showing.
func (m *Map) UpdateMap() {
	var livesMap Cells

	for _, c := range *&m.Cells {
		// checks should c be alive or dead.
		if slices.Contains(m.ToKeepAlive, len(m.Cells.GetLivingNeighbors(c))) {
			livesMap = append(livesMap, c)
		}

		// neighborsCandidates is the dead neighbors of c.
		neighborsCandidates := m.Cells.GetDeadNeighbors(c)

		for _, neighbor := range neighborsCandidates {
			// checks should neighbor become alive or still be dead.
			if slices.Contains(m.ToBecomeAlive, len(m.Cells.GetLivingNeighbors(neighbor))) {
				livesMap = append(livesMap, neighbor)
			}
		}
	}
	livesMap.RemoveDuplicates()

	*&m.Cells = livesMap
}
