package match

import (
	"math"
	"sort"
)

type DistanceBase struct{}

func NewDistanceBase() *DistanceBase {
	return &DistanceBase{}
}

func (d *DistanceBase) Match(t *Individual, list List) {
	temp := make([]*Individual, len(list))
	copy(temp, list)
	sort.Slice(temp, d.Compare(t, list))
	t.SetMatchedIndividual(temp[0])
	temp[0].SetMatchedIndividual(t)
}

func (d *DistanceBase) Compare(target *Individual, list List) func(i, j int) bool {
	// compare with distance
	return func(i, j int) bool {
		tX, tY := target.coord.GetCoord()
		iX, iY := list[i].coord.GetCoord()
		jX, jY := list[j].coord.GetCoord()
		dti := math.Pow(float64(tX-iX), 2) + math.Pow(float64(tY-iY), 2)
		dtj := math.Pow(float64(tX-jX), 2) + math.Pow(float64(tY-jY), 2)
		return math.Sqrt(dti) < math.Sqrt(dtj)
	}
}
