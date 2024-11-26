package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Knubsen/aoc/utils"
)

var contained = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func main() {
	lines, err := utils.ReadLines("")
	utils.Check(err)

	acc := 0
	for _, line := range lines {
		if id, possible := gameIsPossible(line, contained); possible {
			acc += id
		}
	}

	fmt.Println(acc)
}

func gameIsPossible(game string, content map[string]int) (int, bool) {
	prefStructure := strings.Split(game, ":")
	id := strings.Split(prefStructure[0], " ")[1]
	games := strings.Split(prefStructure[1], ";")

	possible := true

	for _, game := range games {
		for _, round := range strings.Split(strings.TrimSpace(game), ",") {
			roundList := strings.Split(strings.TrimSpace(round), " ")

			v := content[roundList[1]]
			cubeCount, err := strconv.Atoi(roundList[0])
			utils.Check(err)

			if cubeCount > v {
				possible = false
			}
		}
	}

	intId, err := strconv.Atoi(id)
	utils.Check(err)

	return intId, possible
}
