package pract4

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	log := NewLog()
	log.LogExecution("Метод Main()")

	op := RunOperation('-', 35)
	op = RunOperation('+', 30)

	_ = op
}
