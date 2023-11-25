package match

type List []*Individual

type MatchMakingSystem struct {
	members     List
	listMap     map[Gender]List
	matchedType MatchType
}

func NewMatchMakingSystem(matchedType MatchType) *MatchMakingSystem {
	return &MatchMakingSystem{
		members:     make(List, 0),
		listMap:     make(map[Gender]List),
		matchedType: matchedType,
	}
}
func (m *MatchMakingSystem) SetMatchedType(matchedType MatchType) {
	m.matchedType = matchedType
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

func (m *MatchMakingSystem) Match(i *Individual) {
	if i.gender == MALE {
		m.matchedType.Match(i, m.GetMatchingList(FEMALE))
	} else {
		m.matchedType.Match(i, m.GetMatchingList(MALE))
	}
}
