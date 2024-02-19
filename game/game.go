package game

import (
	"fmt"
	"github.com/VolodymyrPobochii/game4life/cell"
	"github.com/VolodymyrPobochii/game4life/point"
	"github.com/VolodymyrPobochii/game4life/rule"
	"strings"
)

type Game struct {
	population [][]cell.Cell
	rules      rule.Func
	sb         *strings.Builder
}

func New(rules rule.Func) *Game {
	return &Game{
		rules: rules,
		sb:    &strings.Builder{},
	}
}

func (g *Game) getCell(x, y uint64) *cell.Cell {
	return &g.population[x][y]
}

func (g *Game) gatherNeighbors(c *cell.Cell) {
	maxLen := uint64(len(g.population) - 1)

	pos := c.Pos()

	i, j := pos.X, pos.Y

	if i > 0 && j > 0 {
		tl := g.getCell(i-1, j-1)
		c.AddNeighbor(tl)
	}

	if i > 0 {
		t := g.getCell(i-1, j)
		c.AddNeighbor(t)
	}

	if i > 0 && j < maxLen {
		tr := g.getCell(i-1, j+1)
		c.AddNeighbor(tr)
	}
	if j < maxLen {
		r := g.getCell(i, j+1)
		c.AddNeighbor(r)
	}
	if i < maxLen && j < maxLen {
		br := g.getCell(i+1, j+1)
		c.AddNeighbor(br)
	}

	if i < maxLen {
		b := g.getCell(i+1, j)
		c.AddNeighbor(b)
	}

	if i < maxLen && j > 0 {
		bl := g.getCell(i+1, j-1)
		c.AddNeighbor(bl)
	}

	if j > 0 {
		l := g.getCell(i, j-1)
		c.AddNeighbor(l)
	}

	return
}

func (g *Game) applyRules(c *cell.Cell) {
	g.rules(c)
}

func (g *Game) Init(seed [25][25]uint8) {
	l := len(seed)
	g.population = make([][]cell.Cell, 0, l)
	for i := 0; i < l; i++ {
		g.population = append(g.population, make([]cell.Cell, 0, l))
		for j := 0; j < l; j++ {
			c := cell.New(cell.State(seed[i][j]), point.New(uint64(i), uint64(j)))
			g.population[i] = append(g.population[i], c)
		}
	}
}

func (g *Game) Tick() {
	l := len(g.population)

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			c := &g.population[i][j]
			g.gatherNeighbors(c)
			g.applyRules(c)
		}
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			c := &g.population[i][j]
			c.Transition()
		}
	}
}

func (g *Game) PrintPopulation() {
	g.sb.Reset()

	for _, cells := range g.population {
		for _, c := range cells {
			g.sb.WriteString(c.String())
			g.sb.WriteRune(' ')
		}

		g.sb.WriteRune('\n')
	}

	fmt.Printf("\r%s\n", g.sb.String())
}
