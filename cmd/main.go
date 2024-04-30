package main

import (
	"os"

	svg "github.com/ajstarks/svgo"
	"github.com/dan-frohlich/hexmap"
)

func main() {
	width := 600
	height := 600
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	// canvas.Circle(width/2, height/2, 100)

	h := hexmap.Hexagon{Center: hexmap.Point{X: 300, Y: 300}, Radius: 100}
	xs, ys := convert(h.Vertices())
	canvas.Polygon(xs, ys, `stroke:red`, `stroke-width="1"`)
	canvas.Text(h.Center.X, h.Center.Y, "center", "text-anchor:middle;font-size:30px;fill:white")

	h2 := h.CopyTop()
	xs, ys = convert(h2.Vertices())
	canvas.Polygon(xs, ys, `stroke:red`, `stroke-width="1"`)
	canvas.Text(h2.Center.X, h2.Center.Y, "top", "text-anchor:middle;font-size:30px;fill:white")

	h3 := h.CopyBottom()
	xs, ys = convert(h3.Vertices())
	canvas.Polygon(xs, ys, `stroke:red`, `stroke-width="1"`)
	canvas.Text(h3.Center.X, h3.Center.Y, "bottom", "text-anchor:middle;font-size:30px;fill:white")

	h4 := h.CopyTopLeft()
	xs, ys = convert(h4.Vertices())
	canvas.Polygon(xs, ys, `stroke:red`, `stroke-width="1"`)
	canvas.Text(h4.Center.X, h4.Center.Y, "top left", "text-anchor:middle;font-size:30px;fill:white")

	h5 := h.CopyBottomLeft()
	xs, ys = convert(h5.Vertices())
	canvas.Polygon(xs, ys, `stroke:red`, `stroke-width="1"`)
	canvas.Text(h5.Center.X, h5.Center.Y, "bottom left", "text-anchor:middle;font-size:30px;fill:white")

	h6 := h.CopyTopRight()
	xs, ys = convert(h6.Vertices())
	canvas.Polygon(xs, ys, `stroke:red`, `stroke-width="1"`)
	canvas.Text(h6.Center.X, h6.Center.Y, "top right", "text-anchor:middle;font-size:30px;fill:white")

	h7 := h.CopyBottomRight()
	xs, ys = convert(h7.Vertices())
	canvas.Polygon(xs, ys, `stroke:red`, `stroke-width="1"`)
	canvas.Text(h7.Center.X, h7.Center.Y, "bottom right", "text-anchor:middle;font-size:30px;fill:white")

	canvas.End()
}

func convert(vs []hexmap.Point) (xs []int, ys []int) {
	xs = make([]int, 0, len(vs))
	ys = make([]int, 0, len(vs))
	for _, v := range vs {
		xs = append(xs, v.X)
		ys = append(ys, v.Y)
	}
	return xs, ys
}
