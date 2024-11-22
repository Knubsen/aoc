package main

import (
	"fmt"
	"unicode"

	"github.com/Knubsen/aoc/utils"
)

func main() {
	lines, err := utils.ReadLines("")
	utils.Check(err)

	var acc int
	for _, line := range lines {
		r1, r2 := getOuterDigits(line)
		acc += combineRuneDigits(r1, r2)
	}

	fmt.Println(acc)
}

func combineRuneDigits(r1 rune, r2 rune) int {
	return int(r1-'0')*10 + int(r2-'0')
}

func getOuterDigits(line string) (rune, rune) {
	var first, last rune

	for _, r := range line {
		if unicode.IsDigit(r) {
			if first == 0 {
				first = r
			}
			last = r
		}
	}

	return first, last
}
