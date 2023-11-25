package match

import (
	"sort"
)

type HabbitBase struct{}

func NewHabbitBase() *HabbitBase {
	return &HabbitBase{}
}

func (h *HabbitBase) Match(t *Individual, list List) {
	temp := make([]*Individual, len(list))
	copy(temp, list)
	sort.Slice(temp, h.Compare(t, list))
	t.SetMatchedIndividual(temp[0])
	temp[0].SetMatchedIndividual(t)
}

func (h *HabbitBase) Compare(target *Individual, list List) func(i, j int) bool {
	// compare with habbits
	return func(i, j int) bool {
		targetHabbits := target.habbits.GetHabbits()
		iHabbits := list[i].habbits.GetHabbits()
		jHabbits := list[j].habbits.GetHabbits()
		return h.CountCommonHabbits(targetHabbits, iHabbits) > h.CountCommonHabbits(targetHabbits, jHabbits)
	}
}

func (h *HabbitBase) CountCommonHabbits(h1 Habbits, h2 Habbits) int {
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
