package pract4

import "testing"

func TestTemplateMethod(t *testing.T) {
	val := NewArithmeticProgression(0, 10, 1)
	val.Progress()
	val.Print()

	geom := NewGeometryProgression(1, 10, 2)
	geom.Progress()
	geom.Print()
}
