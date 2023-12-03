package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	aoc.Harness(run)
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
func sumOfCalibrationsValues(input string) int {
	var inputSlice = strings.Split(input, "\n")
	var outputSlice []int
	for _, line := range inputSlice {
		if len(line) == 0 {
			continue
		}
		var lineOutputSlice []int
		for _, char := range line {
			if unicode.IsDigit(char) {
				lineOutputSlice = append(lineOutputSlice, int(char-'0'))
			}
		}
		if len(lineOutputSlice) > 0 {
			outputSlice = append(outputSlice, 10*lineOutputSlice[0]+1*lineOutputSlice[len(lineOutputSlice)-1])
		}
	}
	return sum(outputSlice)
}
func revertSpelledNumbers(word string) string {
	w := word
	spelledNumbers := map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine"}
	for j := 3; j <= len(w); j++ {
		partOfWord := w[0:j]
		for i := 1; i <= 9; i++ {
			if strings.Contains(partOfWord, spelledNumbers[i]) {
				w = strings.Replace(w, spelledNumbers[i], strconv.Itoa(i), 1)
				break
			}
		}
	}
	return w
}
func run(part2 bool, input string) any {
	if part2 {
		return sumOfCalibrationsValues(revertSpelledNumbers(input))
	} else {
		return sumOfCalibrationsValues(input)
	}
}
