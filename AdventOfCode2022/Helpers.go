package AdventOfCode2022

import "os"

func getData(file string) string {
	fileContents, err := os.ReadFile(file)
	check(err)
	return string(fileContents)
}
