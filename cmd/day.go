package cmd

import (
	"fmt"
	"strconv"

	"github.com/Knubsen/aoc/utils"
	"github.com/spf13/cobra"
)

var day = &cobra.Command{
	Use:   "day",
	Short: "Execute day 1 of aoc",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dayNum, err := strconv.Atoi(args[0])
		utils.Check(err)

		part := args[1]
		if part != "p1" && part != "p2" {
			fmt.Println("Invalid part argument. Use 'p1' or 'p2'.")
			return
		}

		dayKey := fmt.Sprintf("d%02d", dayNum)
		funcs, exists := GetDayFuncs(dayKey)
		if !exists {
			fmt.Printf("Day %02d is not implemented.\n", dayNum)
			return
		}

		if part == "p1" {
			funcs.ExecPt01()
		} else if part == "p2" {
			funcs.ExecPt02()
		}
	},
}

func init() {
	rootCmd.AddCommand(day)
}
