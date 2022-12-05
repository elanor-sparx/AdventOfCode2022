package AdventOfCode2022

import (
	"fmt"
	"strings"
)

type Move struct {
	nBoxes         int
	fromStackIndex int
	toStackIndex   int
}

func Day5() {
	data := strings.Split(getData("AdventOfCode2022/input/day5_real.txt"), "\n\n")
	moves := parseMoveData(data[1])
	pic1 := parsePictorialData(data[0])
	pic2 := parsePictorialData(data[0])
	for _, m := range moves {
		pic1 = makeMove(pic1, m)
		pic2 = makeMultiMove(pic2, m)
	}
	fmt.Println(getStackTops(pic1))
	fmt.Println(getStackTops(pic2))
}

func parsePictorialData(data string) []Stack {
	lines := strings.Split(data, "\n")
	nStacks := Max(getIntsFromString(lines[len(lines)-1]))
	stacks := make([]Stack, nStacks)
	for i := len(lines) - 2; i >= 0; i-- {
		runes := parseSinglePicLine(lines[i], nStacks)
		for i, r := range runes {
			if r != 0 {
				stacks[i].Push(r)
			}
		}
	}
	return stacks
}

func parseSinglePicLine(data string, nStacks int) []rune {
	line := []rune(data)
	runes := make([]rune, nStacks)
	for i := 0; i < len(runes); i++ {
		j := 4*i + 1
		if j < len(line) && line[j] != ' ' {
			runes[i] = line[j]
		}
	}
	return runes
}

func parseMoveData(data string) []Move {
	var moves []Move
	for _, line := range strings.Split(data, "\n") {
		m := parseSingleMove(line)
		moves = append(moves, m)
	}
	return moves
}

func parseSingleMove(data string) Move {
	ints := getIntsFromString(data)
	return Move{nBoxes: ints[0], fromStackIndex: ints[1] - 1, toStackIndex: ints[2] - 1}
}

func getStackTops(stacks []Stack) string {
	var tops []rune
	for _, s := range stacks {
		tops = append(tops, s.Peek())
	}
	return string(tops)
}

func makeMove(pic []Stack, m Move) []Stack {
	for i := 0; i < m.nBoxes; i++ {
		box, _ := pic[m.fromStackIndex].Pop()
		pic[m.toStackIndex].Push(box)
	}
	return pic
}

func makeMultiMove(pic []Stack, m Move) []Stack {
	boxes, _ := pic[m.fromStackIndex].PopMulti(m.nBoxes)
	for _, box := range boxes {
		pic[m.toStackIndex].Push(box)
	}
	return pic
}
