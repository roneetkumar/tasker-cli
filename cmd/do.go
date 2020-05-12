package cmd

import (
	"fmt"
	"strconv"

	"github.com/roneetkumar/tasker/db"

	"github.com/spf13/cobra"
)

//DoCMD func
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

		tasks, err := db.AllTasks()

		if err != nil {
			fmt.Println("Something went wrong: ", err.Error())
			return
		}

		for _, id := range ids {

			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}

			task := tasks[id-1]

			err := db.DeleteTask(task.Key)

			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCMD.AddCommand(DoCMD)
}
