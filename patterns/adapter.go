package patterns

import (
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type IGame interface {
	Brosok() int
}

type Kost struct {
}

func (k Kost) Brosok() int {
	return random.Intn(6) + 1
}

type Gamer struct {
	Name string
}

func (g Gamer) String() string {
	return g.Name
}

func (g Gamer) SeansGame(game IGame) int {
	return game.Brosok()
}

type Monet struct {
}

func (m Monet) BrosokM() int {
	return random.Intn(2) + 1
}

type AdapterGame struct {
	mot Monet
}

func (a AdapterGame) Brosok() int {
	return a.mot.BrosokM()
}
