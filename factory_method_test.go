package pract4

import (
	"fmt"
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	trCom := NewTaxiTransCom("Служба такси")
	compService := trCom.Create("Такси", 1)
	dist := 15.5

	PrintFactoryMethodInfo(compService, dist)

	gCom := NewShipTransCom("Служба перевозок")
	compService = gCom.Create("Грузоперевозки", 2)
	distg := 150.5

	PrintFactoryMethodInfo(compService, distg)

}

func PrintFactoryMethodInfo(service TransportService, dist float64) {
	fmt.Printf("Компания %v, расстояние: %v, стоимость: %v\n", service, dist, service.CostTransportation(dist))
}
