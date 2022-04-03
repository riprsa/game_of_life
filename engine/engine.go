package engine

import (
	"golang.org/x/exp/slices"
)

// Cell is a struct that contains the x and y cordinates of a cell.
type Cell struct {
	X int
	Y int
}

// Map is a 2D array of cells that are true. All other cells are false so we don't need to store them.
type Map []Cell

// NewMap returns a new map of cells.
func NewMap() Map {
	m := make(Map, 0)

	// rn i dont have time to rewrite this to normal input or ingame clicks, so i will just hardcode it.
	m = append(m, Cell{X: 0, Y: 0})
	m = append(m, Cell{X: 1, Y: 0})
	m = append(m, Cell{X: 1, Y: 1})
	m = append(m, Cell{X: 2, Y: 1})
	m = append(m, Cell{X: 0, Y: 2})
	return m
}

// GetNeighborsFromMap returns a slice of cells that are adjacent to the cell passed in.
func (m Map) GetNeighborsFromMap(c Cell) []Cell {
	var neighbors []Cell
	for _, cell := range m {
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

// GetCandidateNeighbors returns a slice of cells that are adjacent to the cell passed in and not in Map.
func (m Map) GetCandidateNeighbors(c Cell) []Cell {
	// ALL neighbors of the cell.
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

	// LiveCandidates are the neighbors that MAY BE alive. So it is slice of candidates.
	var сandidates []Cell

	for _, neighbor := range neighbors {
		if slices.Index(m, neighbor) == -1 {
			сandidates = append(сandidates, neighbor)
		}
	}

	return сandidates
}

// RemoveDubles removes duplicates from a Map.
func (m *Map) RemoveDuplicates() {
	var uniqueMap Map
	u := make(map[Cell]struct{})
	for _, c := range *m {
		u[c] = struct{}{}
	}
	for k := range u {
		uniqueMap = append(uniqueMap, k)
	}
	*m = uniqueMap
}

// DeleteCell removes a cell from the map.
func (m *Map) DeleteCell(c Cell) {
	for i, cell := range *m {
		if cell.X == c.X && cell.Y == c.Y {
			*m = append((*m)[:i], (*m)[i+1:]...)
			return
		}
	}
}

// UpdateMap changes the map for one showing.
func (m *Map) UpdateMap() {
	var livesMap Map

	for _, c := range *m {
		neighbors := m.GetNeighborsFromMap(c)

		// checks should the cell die.
		if len(neighbors) == 2 || len(neighbors) == 3 {
			livesMap = append(livesMap, c)
		}

		// neighborsCandidates ARE NOT in Map.
		neighborsCandidates := m.GetCandidateNeighbors(c)

		for _, neighbor := range neighborsCandidates {
			if len(m.GetNeighborsFromMap(neighbor)) == 3 {
				livesMap = append(livesMap, neighbor)
			}
		}
	}
	livesMap.RemoveDuplicates()

	*m = livesMap
}
