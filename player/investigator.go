package player

// Сыщик
type Investigator struct {
	isLead bool
}

func NewInvestigator(isLead bool) *Investigator {
	return &Investigator{
		isLead: isLead,
	}
}
