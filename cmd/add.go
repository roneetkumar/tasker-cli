package cmd

import (
	"github.com/spf13/cobra"
)

//AddCMD
var AddCMD = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
}

func init() {
	RootCMD.AddCommand(AddCMD)
}
