package match

import (
	"sort"
)

type HabbitBaseStrategy struct{}

func NewHabbitBase() *HabbitBaseStrategy {
	return &HabbitBaseStrategy{}
}

func (h *HabbitBaseStrategy) Match(t *Individual, list List) *MatchedIndividual {
	temp := make([]*Individual, len(list))
	copy(temp, list)
	sort.Slice(temp, h.Compare(t, list))
	return NewMatchedIndividual(t, temp[0])
}

func (h *HabbitBaseStrategy) Compare(target *Individual, list List) func(i, j int) bool {
	// compare with habbits
	return func(i, j int) bool {
		targetHabbits := target.habbits.GetHabbits()
		iHabbits := list[i].habbits.GetHabbits()
		jHabbits := list[j].habbits.GetHabbits()
		return h.CountCommonHabbits(targetHabbits, iHabbits) > h.CountCommonHabbits(targetHabbits, jHabbits)
	}
}

func (h *HabbitBaseStrategy) CountCommonHabbits(h1 Habbits, h2 Habbits) int {
	longHabbit := h1
	shortHabbit := h2
	if len(h1) < len(h2) {
		shortHabbit = h1
		longHabbit = h2
	}
	count := 0
	for k := range shortHabbit {
		if _, ok := longHabbit[k]; ok {
			count++
		}
	}
	return count
}
