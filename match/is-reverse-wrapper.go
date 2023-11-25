package match

import "sort"

type IsReverseWrapper[T MatchType] struct {
	origin T
}

func NewIsReverseWarpper[T MatchType](origin T) *IsReverseWrapper[T] {
	return &IsReverseWrapper[T]{
		origin: origin,
	}
}
func (inv *IsReverseWrapper[T]) Match(target *Individual, list List) {
	temp := make([]*Individual, len(list))
	copy(temp, list)
	sort.Slice(temp, inv.Compare(target, list))
	target.SetMatchedIndividual(temp[0])
}

func (inv *IsReverseWrapper[T]) Compare(target *Individual, list List) func(i, j int) bool {
	return func(i int, j int) bool {
		return !inv.origin.Compare(target, list)(i, j)
	}
}
