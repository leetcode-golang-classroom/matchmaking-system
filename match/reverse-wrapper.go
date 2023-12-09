package match

import "sort"

type ReverseWrapper struct {
	matchType MatchType
}

func NewReverseWrapper(matchType MatchType) *ReverseWrapper {
	return &ReverseWrapper{
		matchType: matchType,
	}
}

func (r *ReverseWrapper) Match(target *Individual, list List) *MatchedIndividual {
	temp := make([]*Individual, len(list))
	copy(temp, list)
	sort.Slice(temp, r.Compare(target, list))
	return NewMatchedIndividual(target, temp[0])
}

func (r *ReverseWrapper) Compare(target *Individual, list List) func(i, j int) bool {
	return func(i int, j int) bool {
		return !r.matchType.Compare(target, list)(i, j)
	}
}
