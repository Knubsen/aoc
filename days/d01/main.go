package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Knubsen/aoc/utils"
)

var numberMap = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	lines, err := utils.ReadLines("")
	utils.Check(err)

	var acc int
	for _, line := range lines {
		acc += getFirstDigit(line)*10 + getLastDigit(line)
	}

	fmt.Println(acc)
}

func getLastDigit(line string) int {
	for i := len(line) - 1; i >= 0; i-- {
		if contains, digit := containsSpelledDigit(line[i:]); contains {
			return digit
		} else if unicode.IsDigit(rune(line[i])) {
			return int(rune(line[i]) - '0')
		}
	}

	fmt.Println("No last digit found")
	return 0
}

func getFirstDigit(line string) int {
	for i := 0; i < len(line); i++ {
		if contains, digit := containsSpelledDigit(line[:i]); contains {
			return digit
		} else if unicode.IsDigit(rune(line[i])) {
			return int(rune(line[i]) - '0')
		}
	}

	fmt.Println("No first digit found")
	return 0
}

func containsSpelledDigit(s string) (bool, int) {
	for key, value := range numberMap {
		if strings.Contains(s, key) {
			return true, value
		}
	}
	return false, 0
}
