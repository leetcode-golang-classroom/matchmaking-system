package match

import "sort"

type ReverseWrapper struct {
	matchStrategy MatchStrategy
}

func NewReverseWrapper(matchType MatchStrategy) *ReverseWrapper {
	return &ReverseWrapper{
		matchStrategy: matchType,
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
		return !r.matchStrategy.Compare(target, list)(i, j)
	}
}
