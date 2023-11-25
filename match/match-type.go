package match

type MatchType interface {
	Match(i *Individual, list List)
	Compare(target *Individual, list List) func(i, j int) bool
}
