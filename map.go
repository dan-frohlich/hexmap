package hexmap

import (
	"fmt"
	"os"
	"sort"

	svg "github.com/ajstarks/svgo"
)

type Map struct {
	Width, Height int
	HexRadius     float64
	Units         string
}

func (m Map) Hexes() []Hexagon {
	c := Hexagon{Center: Point{X: int(m.HexRadius + 1), Y: int(m.HexRadius + 1)}, Radius: m.HexRadius}

	traversedCenters := make(map[string]struct{})
	return m.recursiveHexes(c, &traversedCenters)
}

func (m Map) recursiveHexes(root Hexagon, traversedCenters *map[string]struct{}) []Hexagon {
	tc := *traversedCenters
	if _, found := tc[root.Center.String()]; found {
		return nil
	}
	if !m.Contains(root) {
		return nil
	}
	hexes := make([]Hexagon, 0, 7)
	hexes = append(hexes, root)
	tc[root.Center.String()] = struct{}{}

	fns := []func() Hexagon{root.CopyTop, root.CopyTopRight, root.CopyBottomRight, root.CopyBottom, root.CopyBottomLeft, root.CopyTopLeft}

	for _, fn := range fns {
		child := fn()
		hexes = append(hexes, m.recursiveHexes(child, traversedCenters)...)
	}

	return hexes
}

func (m Map) Contains(a Hexagon) bool {
	for _, v := range a.Vertices() {
		if !(v.InBounds(0, 0, m.Width, m.Height)) {
			return false
		}
	}
	return true
}

func (m Map) WriteSVG(f *os.File, points bool, numberHexes bool) {
	canvas := svg.New(f)
	units := m.Units
	if len(units) < 1 {
		units = "mm"
	}
	canvas.StartviewUnit(m.Width, m.Height, units, 0, 0, m.Width, m.Height)

	sz := int(m.HexRadius / 3)
	tyextStyle := fmt.Sprintf("text-anchor:middle;font-family:monospace;font-size:%dpx;fill:blue", sz)
	hexes := m.Hexes()

	sort.Slice(hexes, func(i, j int) bool {
		if hexes[i].Center.X < hexes[j].Center.X {
			return true
		}
		if hexes[i].Center.X > hexes[j].Center.X {
			return false
		}
		if hexes[i].Center.Y < hexes[j].Center.Y {
			return true
		}
		return false
	})

	for i, hex := range hexes {
		vs := hex.Vertices()
		if points {
			for _, v := range vs {
				canvas.Circle(v.X, v.Y, sz/2, `fill:none;stroke:red`, `stroke-width="0.1"`)
			}
		} else {
			xs, ys := Convert(vs)
			canvas.Polygon(xs, ys, `fill:none;stroke:red`, `stroke-width="0.1"`)
		}
		if numberHexes {
			canvas.Text(hex.Center.X, hex.Center.Y+int(m.HexRadius)-sz, fmt.Sprintf("%04d", i), tyextStyle)
		}
	}
	canvas.End()
}
