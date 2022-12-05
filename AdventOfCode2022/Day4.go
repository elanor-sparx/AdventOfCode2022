package AdventOfCode2022

import (
	"fmt"
	"strconv"
	"strings"
)

type Duty struct {
	start int
	stop  int
}

func Day4() {
	data := getData("AdventOfCode2022/input/day4_real.txt")
	workList := strings.Split(data, "\n")
	part1total := 0
	part2total := 0
	for _, roster := range workList {
		f, s := rosterFromString(roster)
		if encompasses(f, s) || encompasses(s, f) {
			part1total++
		}
		if overlaps(f, s) {
			part2total++
		}
	}
	fmt.Println(part1total, part2total)
}

func overlaps(first Duty, second Duty) bool {
	return first.start <= second.stop && first.stop >= second.start
}

func encompasses(first Duty, second Duty) bool {
	return first.start <= second.start && first.stop >= second.stop
}

func rosterFromString(roster string) (Duty, Duty) {
	each := strings.Split(roster, ",")
	return dutyFromString(each[0]), dutyFromString(each[1])
}

func dutyFromString(desc string) Duty {
	hyphenIndex := strings.IndexByte(desc, '-')
	start, _ := strconv.Atoi(desc[:hyphenIndex])
	end, _ := strconv.Atoi(desc[hyphenIndex+1:])
	return Duty{start: start, stop: end}
}
