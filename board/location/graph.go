package location

type Graph struct {
	locations []*Location
	routes    []*Route
}

func NewGraph() *Graph {
	londonLoc := NewLocation("Лондон", City)
	seaLoc := NewLocation("Побережье", Sea)
	wildLoc := NewLocation("Глушь", Wild)

	routes := []*Route{
		0: NewRoute(londonLoc, seaLoc, ShipPath),
		1: NewRoute(londonLoc, wildLoc, TrainPath),
	}

	londonLoc.SetRoutes([]*Route{routes[0], routes[1]})
	seaLoc.SetRoutes([]*Route{routes[0]})
	wildLoc.SetRoutes([]*Route{routes[1]})

	locations := []*Location{
		londonLoc,
		seaLoc,
		wildLoc,
	}

	return &Graph{
		locations: locations,
		routes:    routes,
	}
}

func (graph *Graph) GetList() []*Location {
	return graph.locations
}

func (graph *Graph) GetLocationByName(name string) *Location {
	for _, item := range graph.locations {
		if item.name == name {
			return item
		}
	}

	panic("не найдена локация по имени: " + name)
}
