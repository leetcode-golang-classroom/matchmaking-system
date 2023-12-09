package match

type MatchStrategy interface {
	Match(i *Individual, list List) *MatchedIndividual
	Compare(target *Individual, list List) func(i, j int) bool
}
