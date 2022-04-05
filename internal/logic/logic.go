package logic

import "golang.org/x/exp/slices"

type Scene struct {
	// Map is a 2D array of cells that are true. All other cells are false so we don't need to store them.
	Map Map

	// LivingCells is the living cells.
	LivingCells []Cord

	// ToKeepAlive is a slice of ints that are the number of neighbors that a cell must have to stay alive.
	ToKeepAlive []int

	// ToBecomeAlive is a slice of ints that are the number of neighbors that a cell must have to become alive.
	ToBecomeAlive []int

	IsPaused bool

	Speed int

	Tick int
	tick int
}

func NewScene(toKeepAlive, toBecomeAlive []int) Scene {
	m := NewMap(1000, 1000)

	return Scene{
		Map:           m,
		ToKeepAlive:   toKeepAlive,
		ToBecomeAlive: toBecomeAlive,
		IsPaused:      true,
		Speed:         5, // 1 - 5
		tick:          1,
	}
}

// UpdateScene changes the map for one showing.
func (s *Scene) UpdateScene() {
	if s.IsPaused {
		return
	}

	s.tick++
	if s.tick%(101-s.Speed*20) != 0 {
		return
	}

	newMap := NewMap(len(s.Map), len(s.Map[0]))

	newLivingCells := []Cord{}

	for _, cell := range s.LivingCells {
		if cell.X < 0 || cell.X >= len(s.Map) && cell.Y < 0 || cell.Y >= len(s.Map[0]) {
			continue
		}

		if slices.Contains(s.ToKeepAlive, len(s.Map.GetLivingNeighbors(cell.X, cell.Y))) {
			newMap[cell.X][cell.Y] = true
			newLivingCells = append(newLivingCells, cell)
		}

		neighbors := s.Map.GetDeadNeighbors(cell.X, cell.Y)
		for _, n := range neighbors {
			if slices.Contains(s.ToBecomeAlive, len(s.Map.GetLivingNeighbors(n.X, n.Y))) {
				if n.X < 0 || n.X >= len(newMap) || n.Y < 0 || n.Y >= len(newMap[0]) {
					continue
				}
				newMap[n.X][n.Y] = true
				newLivingCells = append(newLivingCells, n)
			}
		}
	}

	s.LivingCells = removeDuplicates(newLivingCells)

	s.Tick++
	s.Map = newMap
}

func removeDuplicates(s []Cord) []Cord {
	allKeys := make(map[Cord]bool)
	list := []Cord{}
	for _, item := range s {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

type ListOfCords []Cord

type Cord struct {
	X int
	Y int
}

type Map [][]bool

// NewMap returns a new map with the given dimensions. It is filled with false.
func NewMap(width, height int) Map {
	m := make(Map, width)
	for x := range m {
		m[x] = make([]bool, height)
	}

	return m
}

func (m Map) GetLivingNeighbors(x, y int) ListOfCords {
	neighbors := ListOfCords{
		{X: x + 1, Y: y},
		{X: x - 1, Y: y},
		{X: x, Y: y + 1},
		{X: x, Y: y - 1},
		{X: x + 1, Y: y + 1},
		{X: x - 1, Y: y - 1},
		{X: x + 1, Y: y - 1},
		{X: x - 1, Y: y + 1},
	}

	livingNeighbors := ListOfCords{}
	for _, n := range neighbors {
		if n.X < 0 || n.X >= len(m) {
			continue
		}
		if n.Y < 0 || n.Y >= len(m[n.X]) {
			continue
		}
		if m[n.X][n.Y] {
			livingNeighbors = append(livingNeighbors, n)
		}
	}

	return livingNeighbors
}

func (m Map) GetDeadNeighbors(x, y int) ListOfCords {
	neighbors := ListOfCords{
		{X: x + 1, Y: y},
		{X: x - 1, Y: y},
		{X: x, Y: y + 1},
		{X: x, Y: y - 1},
		{X: x + 1, Y: y + 1},
		{X: x - 1, Y: y - 1},
		{X: x + 1, Y: y - 1},
		{X: x - 1, Y: y + 1},
	}

	deads := ListOfCords{}
	for _, n := range neighbors {
		if n.X < 0 || n.X >= len(m) {
			continue
		}
		if n.Y < 0 || n.Y >= len(m[n.X]) {
			continue
		}
		if !m[n.X][n.Y] {
			deads = append(deads, n)
		}
	}

	return deads
}
