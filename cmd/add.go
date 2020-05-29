package cmd

import (
	"fmt"
	"os"
	"strings"

	"gophercises/task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds task(s) to your task list. Use comma to separate multiple tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := strings.Split(strings.Join(args, " "), ",")
		for _, task := range tasks {
			trimmed := strings.TrimSpace(task)
			err := db.CreateTask(trimmed)
			if err != nil {
				fmt.Println("Fail to add task 🥵", err)
				os.Exit(1)
			}
			fmt.Println("Successfully added task 🍀", trimmed)
		}
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
