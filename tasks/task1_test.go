package tasks

import "testing"

func TestAdapter(t *testing.T) {
	system := ClimateControlSystem{sensors: []Thermometer{
		CelsiusThermometer{},
		CelsiusThermometer{},
		FahrenheitThermometerToCelsiusAdapter{FahrenheitThermometer{}},
	}}

	system.PrintMeasurements()
}
