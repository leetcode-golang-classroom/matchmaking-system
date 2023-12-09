package match

import (
	"sort"
)

type DistanceBaseStrategy struct{}

func NewDistanceBase() *DistanceBaseStrategy {
	return &DistanceBaseStrategy{}
}

func (d *DistanceBaseStrategy) Match(t *Individual, list List) *MatchedIndividual {
	temp := make([]*Individual, len(list))
	copy(temp, list)
	sort.Slice(temp, d.Compare(t, list))
	return NewMatchedIndividual(t, temp[0])
}

func (d *DistanceBaseStrategy) Compare(target *Individual, list List) func(i, j int) bool {
	// compare with distance
	return func(i, j int) bool {
		return list[i].coord.Distance(target.coord) < list[j].coord.Distance(target.coord)
	}
}
