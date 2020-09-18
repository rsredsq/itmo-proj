package pract4

import (
	"fmt"
	"testing"
)

func TestAdapter1(t *testing.T) {
	kubik := Kost{}
	g1 := Gamer{Name: "Иван"}
	fmt.Printf("выпало очков %v для игрока %s\n",
		g1.SeansGame(kubik), g1)
}

func TestAdapter2(t *testing.T) {
	g1 := Gamer{Name: "Иван"}
	mon := Monet{}
	var bmon IGame = AdapterGame{mon}
	fmt.Printf("выпало очков %v для игрока %s\n",
		g1.SeansGame(bmon), g1)
}
