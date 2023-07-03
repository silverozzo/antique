package player

// Сыщик
type Investigator struct {
	isLead       bool
	locationName string
	health       int
	sanity       int
}

func NewInvestigator(isLead bool) *Investigator {
	return &Investigator{
		isLead:       isLead,
		locationName: "Лондон",
		health:       5,
		sanity:       5,
	}
}
