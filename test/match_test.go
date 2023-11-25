package test

import (
	"testing"

	"github.com/leetcode-golang-classroom/matchmaking-system/match"
)

func WithMember(tb testing.TB, matchmakingSystem *match.MatchMakingSystem) (*match.Individual, *match.Individual, *match.Individual) {
	john := match.NewIndividual(0, match.MALE, 28, match.NewIntro("My name is John"),
		match.NewHabbits([]string{"maochung", "movie", "music", "guitar", "chess"}), match.NewCoord(1, 4))
	marry := match.NewIndividual(0, match.FEMALE, 32, match.NewIntro("My name is Marry"),
		match.NewHabbits([]string{"tennis", "swimming", "music", "cooking", "reading"}), match.NewCoord(2, 5))
	sue := match.NewIndividual(0, match.FEMALE, 18, match.NewIntro("My name is Sue"),
		match.NewHabbits([]string{"tennis", "movie", "gardening", "reading"}), match.NewCoord(1, 5))
	matchmakingSystem.AddMember(john)
	matchmakingSystem.AddMember(marry)
	matchmakingSystem.AddMember(sue)
	return john, marry, sue
}
func TestDistanceMatchStrategy(t *testing.T) {
	distanceStrategy := match.NewDistanceBase()
	matchmakingSystem := match.NewMatchMakingSystem(distanceStrategy)
	john, _, sue := WithMember(t, matchmakingSystem)
	// testing distance base for john
	matchmakingSystem.Match(john)
	got := john.GetMatched()
	if got.GetValue().GetIntro() != sue.GetIntro() {
		t.Errorf("expected: %s, got: %s", sue.GetIntro(), got.GetValue().GetIntro())
	}
}

func TestHabbitMatchStrategy(t *testing.T) {
	habbitStrategy := match.NewHabbitBase()
	matchmakingSystem := match.NewMatchMakingSystem(habbitStrategy)
	john, marry, _ := WithMember(t, matchmakingSystem)
	// testing habbit base for john
	matchmakingSystem.Match(john)
	got := john.GetMatched()
	if got.GetValue().GetIntro() != marry.GetIntro() {
		t.Errorf("expected: %s, got: %s", marry.GetIntro(), got.GetValue().GetIntro())
	}

}

func TestInverseDistanceStrategy(t *testing.T) {
	distanceStrategy := match.NewDistanceBase()
	// testing is-reverse-wrapper for distance base
	matchmakingSystem := match.NewMatchMakingSystem(match.NewIsReverseWarpper[*match.DistanceBase](distanceStrategy))
	john, marry, _ := WithMember(t, matchmakingSystem)
	matchmakingSystem.Match(john)
	got := john.GetMatched()
	if got.GetValue().GetIntro() != marry.GetIntro() {
		t.Errorf("expected: %s, got: %s", marry.GetIntro(), got.GetValue().GetIntro())
	}
}

func TestInverseHabbitStrategy(t *testing.T) {
	habbitStrategy := match.NewHabbitBase()
	// testing is-reverse-wrapper for habbit base
	matchmakingSystem := match.NewMatchMakingSystem(match.NewIsReverseWarpper[*match.HabbitBase](habbitStrategy))
	john, _, sue := WithMember(t, matchmakingSystem)
	matchmakingSystem.Match(john)
	got := john.GetMatched()
	if got.GetValue().GetIntro() != sue.GetIntro() {
		t.Errorf("expected: %s, got: %s", sue.GetIntro(), got.GetValue().GetIntro())
	}
}
