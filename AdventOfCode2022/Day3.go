package AdventOfCode2022

import (
	"fmt"
	"strings"
	"unicode"
)

func Day3() {
	data := getData("AdventOfCode2022/input/day3_real.txt")
	rucksackList := strings.Split(data, "\n")
	part1total := 0
	part2total := 0
	for i, sack := range rucksackList {
		part1total += getCommonRunePriorityValue(sack)
		if i%3 == 0 {
			part2total += getGroupPriorityValues(rucksackList[i : i+3])
		}
	}
	fmt.Println(part1total, part2total)
}

func getCommonRunePriorityValue(sack string) int {
	left := []rune(sack[:len(sack)/2])
	right := []rune(sack[len(sack)/2:])
	for _, rRune := range right {
		for _, lRune := range left {
			if rRune == lRune {
				return getPriorityValue(rRune)
			}
		}
	}
	return -1
}

func getAllCommonRunes(a []rune, b []rune) []rune {
	var common []rune
	for _, rRune := range a {
		for _, lRune := range b {
			if rRune == lRune {
				common = append(common, rRune)
			}
		}
	}
	return common
}

func getGroupPriorityValues(group []string) int {
	common := getAllCommonRunes([]rune(group[0]), []rune(group[1]))
	return getPriorityValue((getAllCommonRunes(common, []rune(group[2])))[0])
}

func getPriorityValue(r rune) int {
	if unicode.IsLower(r) {
		return int(r-'a') + 1
	} else {
		return int(r-'A') + 27
	}
}
