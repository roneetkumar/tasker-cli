package cmd

import (
	"github.com/spf13/cobra"
)

//RootCMD
var RootCMD = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI task manager",
}
