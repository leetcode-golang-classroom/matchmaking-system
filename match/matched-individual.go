package match

type MatchedIndividual struct {
	matchedIndividual1 *Individual
	matchedIndividual2 *Individual
}

func NewMatchedIndividual(m1, m2 *Individual) *MatchedIndividual {
	return &MatchedIndividual{
		matchedIndividual1: m1,
		matchedIndividual2: m2,
	}
}

func (m *MatchedIndividual) String() string {
	return m.matchedIndividual1.String() + "," + m.matchedIndividual2.String()
}

func (m *MatchedIndividual) GetValue() []*Individual {
	return []*Individual{m.matchedIndividual1, m.matchedIndividual2}
}
