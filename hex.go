package hexmap

import "fmt"

type Point struct {
	X, Y int
}

func (p Point) InBounds(x1, y1, x2, y2 int) bool {
	return p.X >= x1 && p.X <= x2 && p.Y >= y1 && p.Y <= y2
}

func (p Point) Equals(q Point) bool {
	return p.X == q.X && p.Y == q.Y
}

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p Point) Sub(q Point) Point {
	return Point{X: p.X - q.X, Y: p.Y - q.Y}
}

func (p Point) Plus(q Point) Point {
	return Point{X: p.X + q.X, Y: p.Y + q.Y}
}

type Hexagon struct {
	Center Point
	Radius float64
	ID     string
}

const (
	sin60 = 0.8660254037844385965883020617184229195117950439453125000000000000
	cos60 = 0.5
	cos30 = sin60
	sin30 = cos60
)

func (hex Hexagon) CopyTop() Hexagon {
	return Hexagon{Radius: hex.Radius, Center: Point{X: hex.Center.X, Y: -2*int(sin60*hex.Radius) + hex.Center.Y}}
}

func (hex Hexagon) CopyBottom() Hexagon {
	return Hexagon{Radius: hex.Radius, Center: Point{X: hex.Center.X, Y: 2*int(sin60*hex.Radius) + hex.Center.Y}}
}

func (hex Hexagon) CopyTopLeft() Hexagon {
	c2 := hex.a().Sub(hex.Center).Plus(hex.f())
	return Hexagon{Radius: hex.Radius, Center: c2}
}

func (hex Hexagon) CopyBottomLeft() Hexagon {
	c2 := hex.a().Sub(hex.Center).Plus(hex.b())
	return Hexagon{Radius: hex.Radius, Center: c2}
}

func (hex Hexagon) CopyTopRight() Hexagon {
	c2 := hex.e().Sub(hex.Center).Plus(hex.d())
	return Hexagon{Radius: hex.Radius, Center: c2}
}

func (hex Hexagon) CopyBottomRight() Hexagon {
	c2 := hex.d().Sub(hex.Center).Plus(hex.c())
	return Hexagon{Radius: hex.Radius, Center: c2}
}

func (hex Hexagon) BoundingBox() (a, b Point) {
	a = Point{X: (-1 * int(hex.Radius)) + hex.Center.X, Y: -1*int(hex.Radius*sin60) + hex.Center.Y}
	b = Point{X: (int(hex.Radius)) + hex.Center.X, Y: int(hex.Radius*sin60) + hex.Center.Y}
	return a, b
}

func (hex Hexagon) Vertices() []Point {
	/*
	 *    F---E
	 *   /     \
	 *  A   o   D
	 *   \     /
	 *    B---C
	 */

	return []Point{hex.a(), hex.b(), hex.c(), hex.d(), hex.e(), hex.f()}
}
func (hex Hexagon) a() Point {
	return Point{X: (-1 * int(hex.Radius)) + hex.Center.X, Y: hex.Center.Y}
}

func (hex Hexagon) b() Point {
	return Point{X: (-1 * int(hex.Radius*cos60)) + hex.Center.X, Y: int(hex.Radius*sin60) + hex.Center.Y}
}

func (hex Hexagon) c() Point {
	return Point{X: int(hex.Radius*cos60) + hex.Center.X, Y: int(hex.Radius*sin60) + hex.Center.Y}
}

func (hex Hexagon) d() Point {
	return Point{X: int(hex.Radius) + hex.Center.X, Y: hex.Center.Y}
}

func (hex Hexagon) e() Point {
	return Point{X: int(hex.Radius*cos60) + hex.Center.X, Y: -1*int(hex.Radius*sin60) + hex.Center.Y}
}

func (hex Hexagon) f() Point {
	return Point{X: (-1 * int(hex.Radius*cos60)) + hex.Center.X, Y: -1*int(hex.Radius*sin60) + hex.Center.Y}
}

func Convert(vs []Point) (xs []int, ys []int) {
	xs = make([]int, 0, len(vs))
	ys = make([]int, 0, len(vs))
	for _, v := range vs {
		xs = append(xs, v.X)
		ys = append(ys, v.Y)
	}
	return xs, ys
}
