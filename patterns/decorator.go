package patterns

import "fmt"

type AutoBase interface {
	Name() string
	Description() string
	GetCost() float64
}

func AutoBaseToString(a AutoBase) string {
	return fmt.Sprintf("Ваш автомобиль: \n%v \nОписание: %v \nСтоимость %v\n",
		a.Name(), a.Description(), a.GetCost())
}

type Renault struct {
	name        string
	description string
	costBase    float64
}

func (r Renault) Name() string {
	return r.name
}

func (r Renault) Description() string {
	return r.description
}

func (r Renault) GetCost() float64 {
	return r.costBase * 1.18
}

type MediaNav struct {
	AutoBase
	title string
}

func (m MediaNav) Name() string {
	return m.AutoBase.Name() + ". Современный"
}

func (m MediaNav) Description() string {
	return m.AutoBase.Description() + "." + m.title + ". Обновленная мультимедийная навигационная система"
}

func (m MediaNav) GetCost() float64 {
	return m.AutoBase.GetCost() + 15.99
}

type SystemSecurity struct {
	AutoBase
	title string
}

func (s SystemSecurity) Name() string {
	return s.AutoBase.Name() + ". Повышенной безопасности"
}

func (s SystemSecurity) Description() string {
	return s.AutoBase.Description() + "." + s.title + ". Передние боковые безопасности, ESP - система динамической стабилизации автомобиля"
}

func (s SystemSecurity) GetCost() float64 {
	return s.AutoBase.GetCost() + 20.99
}
