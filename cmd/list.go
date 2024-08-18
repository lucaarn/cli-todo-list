/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/lucaarn/cli-todo-list/todo"
	"github.com/mergestat/timediff"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

var allTasks bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your non-completed tasks",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		writer := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		fmt.Fprintln(writer, "ID\tTask\tCreated\tDone")

		todos := todo.Load()
		for _, todo := range todos {
			if !todo.IsComplete || allTasks {
				fmt.Fprintf(writer, "%d\t%s\t%s\t%t\n", todo.ID, todo.Description, timediff.TimeDiff(todo.CreatedAt), todo.IsComplete)
			}
		}

		writer.Flush()
	},
}

func init() {
	listCmd.Flags().BoolVarP(&allTasks, "all", "a", false, "List all tasks (including tasks marked as complete")
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
