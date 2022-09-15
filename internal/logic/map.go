package logic

// Cell is here or not
type Map [][]bool

// Just list of coords, to every use. TODO: rename to make it prettier.
type ListOfCords []Cord

// One point on the map.
type Cord struct {
	X int
	Y int
}

// NewMap returns a new map with the given dimensions. It is filled with false.
func NewMap(width, height int) Map {
	m := make(Map, width)
	for x := range m {
		m[x] = make([]bool, height)
	}

	return m
}

// Get all alive cells around this point.
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

// Get all dead cells around this point.
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
