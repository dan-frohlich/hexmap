package hexmap

import (
	"fmt"
	"math/rand/v2"
	"os"
	"sort"

	svg "github.com/ajstarks/svgo"
)

type Map struct {
	Width, Height int
	HexRadius     float64
	Units         string
}

func (m Map) Hexes(populatePercent int) []*Hexagon {
	c := Hexagon{Center: Point{X: m.Width / 2, Y: m.Height / 2}, Radius: m.HexRadius}
	if populatePercent == 100 || populatePercent == 0 {
		c.Center.X = int(m.HexRadius + 1)
		c.Center.Y = int(m.HexRadius + 1)
	}

	traversedCenters := make(map[string]struct{})
	hexes := m.recursiveHexes(&c, populatePercent, &traversedCenters)
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
	row := 0
	col := 0
	for i, hex := range hexes {
		if i > 0 {
			if hexes[i-1].Center.X < hex.Center.X {
				col++
				row = 0
			} else {
				row++
			}
		}
		hex.ID = fmt.Sprintf("%04d", row+100*col)
	}
	return hexes
}

func (m Map) recursiveHexes(root *Hexagon, populatePercent int, traversedCenters *map[string]struct{}) []*Hexagon {
	tc := *traversedCenters
	if _, found := tc[root.Center.String()]; found {
		return nil
	}
	if !m.Contains(*root) {
		tc[root.Center.String()] = struct{}{}
		return nil
	}
	if rand.IntN(101) > populatePercent {
		tc[root.Center.String()] = struct{}{}
		return nil
	}
	hexes := make([]*Hexagon, 0, 7)
	hexes = append(hexes, root)
	tc[root.Center.String()] = struct{}{}

	fns := []func() Hexagon{root.CopyTop, root.CopyTopRight, root.CopyBottomRight, root.CopyBottom, root.CopyBottomLeft, root.CopyTopLeft}

	for _, fn := range fns {
		child := fn()
		hexes = append(hexes, m.recursiveHexes(&child, populatePercent, traversedCenters)...)
	}

	return hexes
}

func (m Map) Contains(a Hexagon) bool {
	v1, v2 := a.BoundingBox()
	return v1.InBounds(0, 0, m.Width, m.Height) && v2.InBounds(0, 0, m.Width, m.Height)
}

func (m Map) WriteSVG(f *os.File, points bool, numberHexes bool, populatePercent int) {
	canvas := svg.New(f)
	units := m.Units
	if len(units) < 1 {
		units = "mm"
	}
	canvas.StartviewUnit(m.Width, m.Height, units, 0, 0, m.Width, m.Height)

	sz := int(m.HexRadius / 3)
	tyextStyle := fmt.Sprintf("text-anchor:middle;font-family:monospace;font-size:%dpx;fill:blue", sz)
	hexes := m.Hexes(populatePercent)

	for _, hex := range hexes {
		vs := hex.Vertices()
		if points {
			for _, v := range vs {
				canvas.Circle(v.X, v.Y, sz/2, `fill:none;stroke:red`, `stroke-width="0.1"`)
			}
		} else {
			xs, ys := Convert(vs)
			canvas.Polygon(xs, ys, `fill:black;stroke:red`, `stroke-width="0.1"`)
		}
		if numberHexes {
			canvas.Text(hex.Center.X, hex.Center.Y+int(m.HexRadius)-sz, hex.ID, tyextStyle)
		}
	}
	canvas.End()
}
