package pract4

import (
	"fmt"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	factory := FordFactory{}

	client := NewClient(factory)

	fmt.Printf("Максимальная скорость %v составляет %v км/час\n", client.String(), client.RunMaxSpeed())
}
