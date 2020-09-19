package tasks

import "fmt"

type RouteStrategy interface {
	BuildRoute()
}

type AutoRoad struct {
}

func (r AutoRoad) BuildRoute() {
	fmt.Println("построен маршрут по автодороге")
}

type WalkRoad struct {
}

func (r WalkRoad) BuildRoute() {
	fmt.Println("построен маршрут по пешей дороге")
}

type BicycleRoad struct {
}

func (r BicycleRoad) BuildRoute() {
	fmt.Println("построен маршрут по велодорожке")
}

type PublicTransportRoad struct {
}

func (r PublicTransportRoad) BuildRoute() {
	fmt.Println("построен маршрут с использованием общественного транспорта")
}

type AttractionsRoad struct {
}

func (r AttractionsRoad) BuildRoute() {
	fmt.Println("построен маршрут для посещения достопримечательностей")
}

type Navigator struct {
	mapShown      bool
	routeStrategy RouteStrategy
}

func (n *Navigator) ShowMap() {
	n.mapShown = true
}

func (n *Navigator) SetRouteStrategy(strategy RouteStrategy) {
	n.routeStrategy = strategy
}

func (n *Navigator) BuildRoute() {
	if n.routeStrategy != nil {
		n.routeStrategy.BuildRoute()
	}
}
