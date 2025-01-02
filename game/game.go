package game

import (
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 500
	screenHeight = 400
	cellSize     = 10
	width        = screenWidth / cellSize
	height       = screenHeight / cellSize
)

type Game struct {
	grid Grid
}

type Grid [][]bool

func NewGrid() Grid {
	grid := make(Grid, height)
	for i := range grid {
		grid[i] = make([]bool, width)
	}
	return grid
}

func NewRandomGrid() Grid {
	grid := NewGrid()
	rand.Seed(time.Now().UnixNano())
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = rand.Float64() < 0.5
		}
	}
	return grid
}

func (g Grid) Next() Grid {
	newGrid := NewGrid()
	for y := range g {
		for x := range g[y] {
			aliveNeighbors := g.aliveNeighbors(x, y)
			if g[y][x] {
				newGrid[y][x] = aliveNeighbors == 2 || aliveNeighbors == 3
			} else {
				newGrid[y][x] = aliveNeighbors == 3
			}
		}
	}
	return newGrid
}

func (g Grid) aliveNeighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			nx, ny := (x+dx+width)%width, (y+dy+height)%height
			if g[ny][nx] {
				count++
			}
		}
	}
	return count
}

func (g *Game) Update() error {
	g.grid = g.grid.Next()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	for y := range g.grid {
		for x := range g.grid[y] {
			if g.grid[y][x] {
				ebitenutil.DrawRect(screen, float64(x*cellSize), float64(y*cellSize), cellSize, cellSize, color.White)
			}
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func RunGame() {
	grid := NewRandomGrid()
	// Initialize with a simple pattern (glider)

	game := &Game{grid: grid}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Conway's Game of Life")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
