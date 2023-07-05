package location

type PathType int

const (
	UnchartedPath PathType = iota
	TrainPath
	ShipPath
)

type Route struct {
	locations [2]*Location
	path      PathType
}

func NewRoute(loc1 *Location, loc2 *Location, path PathType) *Route {
	return &Route{
		locations: [2]*Location{loc1, loc2},
		path:      path,
	}
}

func (rt *Route) GetAnotherEnd(loc *Location) *Location {
	if rt.locations[0] == loc {
		return rt.locations[1]
	}

	if rt.locations[1] == loc {
		return rt.locations[0]
	}

	panic("в текущем пути нет указанной примыкающей локации")
}
