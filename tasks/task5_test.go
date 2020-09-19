package tasks

import "testing"

func TestNavigator(t *testing.T) {
	n := Navigator{}
	n.ShowMap()

	n.SetRouteStrategy(AutoRoad{})
	n.BuildRoute()

	n.SetRouteStrategy(WalkRoad{})
	n.BuildRoute()

	n.SetRouteStrategy(PublicTransportRoad{})
	n.BuildRoute()

	n.SetRouteStrategy(BicycleRoad{})
	n.BuildRoute()

	n.SetRouteStrategy(AttractionsRoad{})
	n.BuildRoute()
}
