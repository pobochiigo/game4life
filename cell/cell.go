package cell

import (
	"fmt"
	"github.com/VolodymyrPobochii/game4life/point"
)

type State uint8

const (
	StateDead State = iota
	StateAlive
)

type Cell struct {
	state     State
	nextState State
	pos       point.Point
	n         uint8
}

func New(state State, pos point.Point) Cell {
	return Cell{
		state: state,
		pos:   pos,
	}
}

func (c *Cell) Pos() point.Point {
	return c.pos
}

func (c *Cell) Transition() {
	c.state = c.nextState
	c.n = 0
}

func (c *Cell) DieNext() {
	c.nextState = StateDead
}

func (c *Cell) LiveNext() {
	c.nextState = StateAlive
}

func (c *Cell) IsAlive() bool {
	return c.state == StateAlive
}

func (c *Cell) AddNeighbor(nc *Cell) {
	c.n += uint8(nc.state)
}

func (c *Cell) Neighbors() uint8 {
	return c.n
}

func (c *Cell) String() string {
	return fmt.Sprintf("%d", c.state)
}
