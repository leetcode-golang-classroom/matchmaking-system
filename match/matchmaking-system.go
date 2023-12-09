package match

type List []*Individual

type MatchMakingSystem struct {
	members         List
	listMap         map[Gender]List
	matchedStrategy MatchStrategy
}

func NewMatchMakingSystem(matchedType MatchStrategy) *MatchMakingSystem {
	return &MatchMakingSystem{
		members:         make(List, 0),
		listMap:         make(map[Gender]List),
		matchedStrategy: matchedType,
	}
}
func (m *MatchMakingSystem) SetMatchedType(matchedType MatchStrategy) {
	m.matchedStrategy = matchedType
}
func (m *MatchMakingSystem) GetCurrentMembersLength() int {
	return len(m.members)
}
func (m *MatchMakingSystem) AddMember(i *Individual) {
	i.id = len(m.members) + 1
	m.members = append(m.members, i)
	m.listMap[i.gender] = append(m.listMap[i.gender], i)
}
func (m *MatchMakingSystem) GetMembers() List {
	return m.members
}
func (m *MatchMakingSystem) GetMatchingList(gender Gender) List {
	return m.listMap[gender]
}

func (m *MatchMakingSystem) Match(i *Individual) *MatchedIndividual {
	if i.gender == MALE {
		return m.matchedStrategy.Match(i, m.GetMatchingList(FEMALE))
	} else {
		return m.matchedStrategy.Match(i, m.GetMatchingList(MALE))
	}
}
