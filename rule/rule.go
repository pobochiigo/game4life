package rule

import "github.com/VolodymyrPobochii/game4life/cell"

type Func func(c *cell.Cell)

// Underpopulation Any live cell with fewer than two live neighbors
// dies as if caused by underpopulation
func Underpopulation(next Func) Func {
	return func(c *cell.Cell) {
		if c.IsAlive() && c.Neighbors() < 2 {
			c.DieNext()
			return
		}

		if next != nil {
			next(c)
		}
	}
}

// Normal Any live cell with two or three live neighbors
// lives on to the next generation
func Normal(next Func) Func {
	return func(c *cell.Cell) {
		n := c.Neighbors()

		if c.IsAlive() && n > 1 && n < 4 {
			c.LiveNext()
			return
		}

		if next != nil {
			next(c)
		}
	}
}

// Overcrowding Any live cell with more than three live neighbors
// dies as if by overcrowding
func Overcrowding(next Func) Func {
	return func(c *cell.Cell) {
		if c.IsAlive() && c.Neighbors() > 3 {
			c.DieNext()
			return
		}

		if next != nil {
			next(c)
		}
	}
}

// Reproduction Any dead cell with exactly three live neighbors
// becomes a live cell, as if by reproduction
func Reproduction(next Func) Func {
	return func(c *cell.Cell) {
		if !c.IsAlive() && c.Neighbors() == 3 {
			c.LiveNext()
			return
		}

		if next != nil {
			next(c)
		}
	}
}
