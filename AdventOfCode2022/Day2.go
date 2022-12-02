package AdventOfCode2022

import (
	"fmt"
	"strings"
)

type singlePlay struct {
	value   int
	beats   int
	losesto int
}

var allPlays = map[int]singlePlay{}
var scores = map[string]int{}

func init() {
	allPlays[0] = singlePlay{value: 1, beats: 3, losesto: 2} //rock
	allPlays[1] = singlePlay{value: 2, beats: 1, losesto: 3} //paper
	allPlays[2] = singlePlay{value: 3, beats: 2, losesto: 1} //scissors

	scores["win"] = 6
	scores["draw"] = 3
	scores["lose"] = 0
}

func Day2() {
	data := getData("AdventOfCode2022/input/day2_real.txt")
	strategyList := strings.Split(data, "\n")
	part1total := 0
	part2total := 0
	for _, round := range strategyList {
		runes := []rune(round)
		part1total += score(allPlays[int(runes[0]-'A')], allPlays[int(runes[2]-'X')])
		part2total += fix(allPlays[int(runes[0]-'A')], runes[2])
	}
	fmt.Println(part1total, part2total)
}

func score(oppositionPlay singlePlay, myPlay singlePlay) int {
	if myPlay.beats == oppositionPlay.value {
		return scores["win"] + myPlay.value
	} else if myPlay.value == oppositionPlay.value {
		return scores["draw"] + myPlay.value
	} else if myPlay.value == oppositionPlay.beats {
		return scores["lose"] + myPlay.value
	} else {
		return -1 // this should not be possible
	}
}

func fix(oppositionPlay singlePlay, myPlayStyle rune) int {
	switch myPlayStyle {
	case 'X': // lose
		return allPlays[oppositionPlay.beats-1].value + scores["lose"]
	case 'Y': // draw
		return oppositionPlay.value + scores["draw"]
	case 'Z': // win
		return allPlays[oppositionPlay.losesto-1].value + scores["win"]
	default:
		return -1 // should not be possible
	}
}
