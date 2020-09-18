package pract4

import "fmt"

type TransportService interface {
	CostTransportation(distance float64) float64
	fmt.Stringer
}

type TransportCompany interface {
	Create(n string, k int) TransportService
}

type TaxiServices struct {
	name     string
	category int
}

func NewTaxiServices(name string, category int) TransportService {
	return TaxiServices{name, category}
}

func (t TaxiServices) CostTransportation(distance float64) float64 {
	return distance * float64(t.category)
}

func (t TaxiServices) String() string {
	return fmt.Sprintf("Фирма %v, поездка категории %v", t.name, t.category)
}

type Shipping struct {
	name   string
	tariff float64
}

func NewShipping(name string, tariff float64) TransportService {
	return Shipping{name, tariff}
}

func (s Shipping) CostTransportation(distance float64) float64 {
	return distance * s.tariff
}

func (s Shipping) String() string {
	return fmt.Sprintf("Фирма %v, доставка по тарифу %v", s.name, s.tariff)
}

type TaxiTransCom struct {
	name string
}

func NewTaxiTransCom(name string) TransportCompany {
	return TaxiTransCom{name}
}

func (t TaxiTransCom) Create(n string, k int) TransportService {
	return NewTaxiServices(n, k)
}

type ShipTransCom struct {
	name string
}

func NewShipTransCom(name string) TransportCompany {
	return ShipTransCom{name}
}

func (s ShipTransCom) Create(n string, k int) TransportService {
	return NewShipping(n, float64(k))
}
