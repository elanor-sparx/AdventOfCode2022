package AdventOfCode2022

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Day1() {
	for _, n := range []int{1, 3} {
		answer, err := sumTopN(caloriesInGroup(), n)
		check(err)
		fmt.Println(answer)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCalorieGroups(data string) []string {
	return strings.Split(data, "\n\n")
}

func sumCaloriesInGroup(data string) int {
	foods := strings.Split(data, "\n")
	calories := 0
	for i := range foods {
		c, err := strconv.Atoi(foods[i])
		check(err)
		calories += c
	}
	return calories
}

func sumTopN(array []int, n int) (int, error) {
	if n > len(array) {
		return -1, errors.New(fmt.Sprintf("There are fewer than %d elements in the array", n))
	}
	ret := 0
	for i := 0; i < n; i++ {
		ret += array[i]
	}
	return ret, nil
}

func caloriesInGroup() []int {
	foodGroups := getCalorieGroups(getData("AdventOfCode2022/input/day1_real.txt"))
	var allCalories []int
	for i := range foodGroups {
		allCalories = append(allCalories, sumCaloriesInGroup(foodGroups[i]))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(allCalories)))
	return allCalories
}
