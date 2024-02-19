package point

type Point struct {
	X, Y uint64
}

func New(x, y uint64) Point {
	return Point{x, y}
}
