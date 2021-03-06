package commands

import (
	"github.com/tomekbielaszewski/goMultiTimer/src/internal/app/gotimer/task"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start task",
	Run: func(cmd *cobra.Command, args []string) {
		taskName := args[0]
		task.NewPomodoro(taskName, 500, 500, 500, 4).Start()
	},
}
