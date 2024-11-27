package day2

import (
	"fmt"

	"github.com/Knubsen/aoc/utils"
)

type game struct {
	red   int
	green int
	blue  int
}

func ExecPt02() {
	lines, err := utils.ReadLines("d02")
	utils.Check(err)

	acc := 0
	for _, line := range lines {
		g := smallestGame(line)
		acc += g.red * g.green * g.blue
	}

	fmt.Println(acc)
}

func smallestGame(line string) *game {
	var currentGame game

	return nil
}
