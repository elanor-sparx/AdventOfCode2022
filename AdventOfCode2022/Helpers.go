package AdventOfCode2022

import (
	"os"
	"regexp"
	"strconv"
)

func getData(file string) string {
	fileContents, err := os.ReadFile(file)
	check(err)
	return string(fileContents)
}

func getIntsFromString(data string) []int {
	var ints []int
	re := regexp.MustCompile("[0-9]+")
	intStrs := re.FindAllString(data, -1)
	for _, s := range intStrs {
		i, _ := strconv.Atoi(s)
		ints = append(ints, i)
	}
	return ints
}

func Max(ints []int) int {
	max := ints[0]
	for _, i := range ints[1:len(ints)] {
		if i > max {
			max = i
		}
	}
	return max
}
