package patterns

import "fmt"

type CarFactory interface {
	CreateCar() AbstractCar
	CreateEngine() AbstractEngine
}

type AbstractCar interface {
	fmt.Stringer
	Name() string
	MaxSpeed(AbstractEngine) int
}

type AbstractEngine interface {
	MaxSpeed() int
}

type FordFactory struct {
}

func (f FordFactory) CreateCar() AbstractCar {
	return FordCar{"Форд"}
}

func (f FordFactory) CreateEngine() AbstractEngine {
	return FordEngine{}
}

type FordCar struct {
	name string
}

func (f FordCar) Name() string {
	return f.name
}

func (f FordCar) MaxSpeed(e AbstractEngine) int {
	return e.MaxSpeed()
}

func (f FordCar) String() string {
	return "Автомобиль " + f.name
}

type FordEngine struct {
}

func (f FordEngine) MaxSpeed() int {
	return 220
}

type Client struct {
	car    AbstractCar
	engine AbstractEngine
}

func NewClient(factory CarFactory) Client {
	return Client{
		car:    factory.CreateCar(),
		engine: factory.CreateEngine(),
	}
}

func (c Client) RunMaxSpeed() int {
	return c.car.MaxSpeed(c.engine)
}

func (c Client) String() string {
	return c.car.String()
}
