package main

import (
	"fmt"

	"github.com/Knubsen/aoc/utils"
)

func main() {
	lines, err := utils.ReadLines("")
	utils.Check(err)

	fmt.Println(lines)
}
