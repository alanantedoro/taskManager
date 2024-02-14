package cmd

import (
	"fmt"
	"os"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your todo list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong", err.Error())
			os.Exit(1)
		}
		fmt.Printf("added \"%s\" to your task list", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
