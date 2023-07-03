package location

type Graph struct {
	locations []*Location
}

func NewGraph() *Graph {
	londonLoc := NewLocation("Лондон", City)
	oceanLoc := NewLocation("Побережье", Ocean)
	wildLoc := NewLocation("Глушь", Wild)

	londonLoc.SetNeigbors([]*Location{oceanLoc, wildLoc})
	oceanLoc.SetNeigbors([]*Location{londonLoc})
	wildLoc.SetNeigbors([]*Location{londonLoc})

	locations := [3]*Location{
		londonLoc,
		oceanLoc,
		wildLoc,
	}

	return &Graph{
		locations: locations[:],
	}
}

func (graph *Graph) GetLocationByName(name string) *Location {
	for _, item := range graph.locations {
		if item.name == name {
			return item
		}
	}

	panic("не найдена локация по имени: " + name)
}
