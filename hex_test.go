package hexmap

import (
	"testing"
)

func Test_HexVertices(t *testing.T) {
	h := Hexagon{Center: Point{100, 150}, Radius: 50}

	vs := h.Vertices()
	if len(vs) != 6 {
		t.Fatalf("only %d vertices found in heagon", len(vs))
	}

	for i, v := range vs {
		t.Logf("vertex: %s\n", v)

		if i > 0 {
			for j := 0; j < i; j++ {
				if vs[j].Equals(v) {
					t.Fatalf("vertex at %d [%s] equals vertex at %d [%s]", i, v, j, vs[j])
				}
			}
		}
	}

	h2 := h.CopyTopRight()
	vs = h2.Vertices()
	for _, v := range vs {
		t.Logf("cp vertex: %s\n", v)
	}
	// t.Logf("sin(60): %.64f", math.Sqrt(3)/2)

	// t.Logf("cos(30): %.64f", math.Cos(math.Pi/6))
}
