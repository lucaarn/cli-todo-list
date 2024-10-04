package cmd

import (
	"github.com/lucaarn/cli-todo-list/todo"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "deletes task by given id",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln("id is not a number", err)
		}

		var todos todo.TodoList
		for _, todo := range todo.Load() {
			if todo.ID != id {
				todos = append(todos, todo)
			}
		}

		todos.Save()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
