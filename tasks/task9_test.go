package tasks

import (
	"testing"
)

func TestCommand(t *testing.T) {
	calculator := NewCalculator()
	result := 0.0
	result = calculator.Add(5)
	if result != 5 {
		t.Fatal()
	}
	result = calculator.Add(4)
	if result != 9 {
		t.Fatal()
	}
	result = calculator.Add(3)
	if result != 12 {
		t.Fatal()
	}
	result = calculator.Redo()
	if result != 15 {
		t.Fatal()
	}
	result = calculator.Undo()
	if result != 12 {
		t.Fatal()
	}
	result = calculator.Sub(2)
	if result != 10 {
		t.Fatal()
	}
	result = calculator.Redo()
	if result != 8 {
		t.Fatal()
	}
	result = calculator.Undo()
	if result != 10 {
		t.Fatal()
	}
	result = calculator.Undo()
	if result != 12 {
		t.Fatal()
	}
	result = calculator.Mul(2)
	if result != 24 {
		t.Fatal()
	}
	result = calculator.Redo()
	if result != 48 {
		t.Fatal()
	}
	result = calculator.Undo()
	if result != 24 {
		t.Fatal()
	}
	result = calculator.Div(2)
	if result != 12 {
		t.Fatal()
	}
	result = calculator.Redo()
	if result != 6 {
		t.Fatal()
	}
	result = calculator.Undo()
	if result != 12 {
		t.Fatal()
	}
	result = calculator.Mul(2)
	if result != 24 {
		t.Fatal()
	}
	result = calculator.Add(2)
	if result != 26 {
		t.Fatal()
	}
	result = calculator.RedoN(2)
	if result != 54 {
		t.Fatal()
	}
	result = calculator.UndoN(2)
	if result != 26 {
		t.Fatal()
	}
}
