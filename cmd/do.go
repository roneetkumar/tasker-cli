package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

//DoCMD
var DoCMD = &cobra.Command{
	Use:   "do",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {

		var ids []int

		for _, arg := range args {

			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			}
			ids = append(ids, id)
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCMD.AddCommand(DoCMD)
}
