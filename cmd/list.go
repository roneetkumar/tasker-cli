package cmd

import (
	"github.com/spf13/cobra"
)

// ListCMD
var ListCMD = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",
}

func init() {
	RootCMD.AddCommand(AddCMD)
}
