package patterns

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {
	calculator := NewCalculator()
	result := 0.0
	result = calculator.Add(5)
	fmt.Println(result)
	result = calculator.Add(4)
	fmt.Println(result)
	result = calculator.Add(3)
	fmt.Println(result)
	result = calculator.Redo()
	fmt.Println(result)
	result = calculator.Undo()
	fmt.Println(result)
}
