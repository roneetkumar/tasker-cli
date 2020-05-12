package cmd

import (
	"github.com/spf13/cobra"
)

//RootCMD struct
var RootCMD = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
