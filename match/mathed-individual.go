package match

type MatchedIndividual struct {
	matched *Individual
}

func NewMatchedIndividual(i *Individual) *MatchedIndividual {
	return &MatchedIndividual{
		matched: i,
	}
}

func (m *MatchedIndividual) String() string {
	return m.matched.String()
}

func (m *MatchedIndividual) GetValue() *Individual {
	return m.matched
}
