package AdventOfCode2022

import (
	"fmt"
)

func Day6() {
	data := []rune(getData("AdventOfCode2022/input/day6_real.txt"))
	fmt.Println(markerCount(data, 4))
	fmt.Println(markerCount(data, 14))
	//test()
}

func test() {
	all_tests := []string{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", "bvwbjplbgvbhsrlpgdmjqwftvncz", "nppdvjthqldpwncqszvftbrmjlhg", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}
	for i, data := range all_tests {
		d := []rune(data)
		fmt.Println(i, markerCount(d, 4))
		fmt.Println(i, markerCount(d, 14))
	}
}

func markerCount(data []rune, n int) int {
	for i := n; i < len(data); i++ {
		if isSet(data[i-n : i]) {
			return i
		}
	}
	return -1
}

func isIn(r rune, data []rune) bool {
	for _, s := range data {
		if r == s {
			return true
		}
	}
	return false
}

func isSet(data []rune) bool {
	var set []rune
	set = append(set, data[0])
	for _, s := range data[1:] {
		if isIn(s, set) {
			return false
		}
		set = append(set, s)
	}
	return len(set) == len(data)
}
