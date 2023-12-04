package main

import (
	"fmt"

	"github.com/leetcode-golang-classroom/matchmaking-system/match"
)

func main() {
	// initial with strategy
	distanceStrategy := match.NewDistanceBase()
	habbitStrategy := match.NewHabbitBase()
	matchingSystem := match.NewMatchMakingSystem(distanceStrategy)
	// setup members
	john := match.NewIndividual(0, match.MALE, 28, match.NewIntro("My name is John"),
		match.NewHabbits([]string{"maochung", "movie", "music", "guitar", "chess"}), match.NewCoord(1, 4))
	marry := match.NewIndividual(0, match.FEMALE, 32, match.NewIntro("My name is Marry"),
		match.NewHabbits([]string{"tennis", "swimming", "music", "cooking", "reading"}), match.NewCoord(2, 5))
	sue := match.NewIndividual(0, match.FEMALE, 18, match.NewIntro("My name is Sue"),
		match.NewHabbits([]string{"tennis", "movie", "gardening", "reading"}), match.NewCoord(1, 5))
	matchingSystem.AddMember(john)
	matchingSystem.AddMember(marry)
	matchingSystem.AddMember(sue)
	// match with distance stratgey
	matchingSystem.Match(john)
	// clear result
	// switch to habbitBase strategy
	matchingSystem.SetMatchedType(habbitStrategy)
	matchedResult := matchingSystem.Match(john)
	fmt.Println(matchedResult) // marry is habbit matched
	// switch to distance reverse stratgey
	matchingSystem.SetMatchedType(match.NewIsReverseWarpper[*match.DistanceBase](distanceStrategy))
	matchedResult = matchingSystem.Match(john)
	fmt.Println(matchedResult) // marry is longest distance for matched
	matchingSystem.SetMatchedType(match.NewIsReverseWarpper[*match.HabbitBase](habbitStrategy))
	matchedResult = matchingSystem.Match(john)
	fmt.Println(matchedResult) // sue is longest distance for matched
}
