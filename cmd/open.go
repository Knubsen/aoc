package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens notes cli",
	Run: func(cmd *cobra.Command, args []string) {

		// err := startREPL(cmd)
		// if err != nil {
		// 	fmt.Println("Error:", err)
		// }

		fmt.Println("Opened aoc cli")

	},
}
