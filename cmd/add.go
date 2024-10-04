package cmd

import (
	"github.com/lucaarn/cli-todo-list/todo"
	"github.com/spf13/cobra"
	"time"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <description>",
	Short: "Add a task to the list with the given description",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		todos := todo.Load()
		todos = append(todos, todo.Todo{
			ID:          todo.GetLatestID(todos) + 1,
			Description: description,
			CreatedAt:   time.Now(),
			IsComplete:  false,
		})
		todos.Save()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
