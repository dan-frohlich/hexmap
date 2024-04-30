package hexmap

import (
	"os"
	"testing"
)

func Test_map(t *testing.T) {

	//us letter is 215.9 x 279.4 mm
	// subtract -12.7mm per dimension for a 1/4" (6.3mm) gutter
	var letterWidth float64 = 215.9
	var letterHeight float64 = 279.4
	m := Map{Width: int(letterWidth - 12.7), Height: int(letterHeight - 12.7), HexRadius: 12.7, Units: "mm"}
	hexes := m.Hexes()

	t.Logf("found %d hexes", len(hexes))

	f, err := os.Create(".map_test.svg")
	if err != nil {
		t.Fatal(err)
	}
	m.WriteSVG(f, false, true)

	f, err = os.Create(".map_test_2.svg")
	if err != nil {
		t.Fatal(err)
	}
	m.WriteSVG(f, true, false)
}
