package tasks

import (
	"fmt"
	"math/rand"
	"time"
)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

type Thermometer interface {
	// returns temperature in celsius
	CaptureTemperature() float64
}

type ClimateControlSystem struct {
	sensors []Thermometer
}

func (c *ClimateControlSystem) PrintMeasurements() {
	for i, sensor := range c.sensors {
		fmt.Printf("температура датчика №%v: %v °C\n", i, sensor.CaptureTemperature())
	}
}

type CelsiusThermometer struct {
}

func (t CelsiusThermometer) CaptureTemperature() float64 {
	return random.Float64() * 25
}

type FahrenheitThermometer struct {
}

func (t FahrenheitThermometer) MakeTemperatureMeasurementInFahrenheit() int {
	return random.Intn(228)
}

type FahrenheitThermometerToCelsiusAdapter struct {
	FahrenheitThermometer
}

func (t FahrenheitThermometerToCelsiusAdapter) CaptureTemperature() float64 {
	return float64(t.FahrenheitThermometer.MakeTemperatureMeasurementInFahrenheit()-32) * 5.0 / 9.0
}
