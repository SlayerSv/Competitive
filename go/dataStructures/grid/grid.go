package grid

import "fmt"

type Grid struct {
	n int
	m int
	grid [][]byte
}

type Pt struct {
	R, C int
}

func (p Pt) Add(o Pt) Pt { return Pt{p.R + o.R, p.C + o.C} }

var Dirs4 = []Pt{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
var Dirs8 = []Pt{{-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}, {-1, -1}}

func NewGrid(n, m int) *Grid {
	grid := make([][]byte, n)
	for i := 0; i < n; i++ {
		grid[i] = nextLine()
	}
	return &Grid{n, m, grid}
}

func NewGridPad(n, m int, char byte) *Grid {
	grid := make([][]byte, n+2)
	grid[0] = make([]byte, m+2)
	grid[n+1] = make([]byte, m+2)
	if char != 0 {
		for i := 0; i <= m+1; i++ {
			grid[0][i] = char
			grid[n+1][i] = char
		}
	}
	for i := 1; i <= n; i++ {
		row := make([]byte, m+2)
		line := nextLineSlice()
		copy(row[1:m+1], line)
		row[0], row[m+1] = char, char
		grid[i] = row
	}
	return &Grid{n, m, grid}
}

func (g *Grid) Find(char byte) Pt {
	for r := 0; r < g.n; r++ {
		for c := 0; c < g.m; c++ {
			if g.grid[r][c] == char {
				return Pt{r, c}
			}
		}
	}
	return Pt{-1, -1}
}

func (g *Grid) Copy() *Grid {
	newG := make([][]byte, g.n)
	for i := 0; i < g.n; i++ {
		newG[i] = make([]byte, g.m)
		copy(newG[i], g.grid[i])
	}
	return &Grid{g.n, g.m, newG}
}

func (g *Grid) Print() {
	for _, row := range g.grid {
		fmt.Printf("%s\n", row)
	}
}


func nextLine() []byte {
	return nil
}

func nextLineSlice() []byte {
	return nil
}