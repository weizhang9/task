package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gophercises/task/db"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Makrs task(s) as done. Use comma or space to separate multiple tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("You have no tasks to delete 🥶")
			os.Exit(1)
		}

		var ids []string
		for _, arg := range args {
			trimmed := strings.Trim(arg, ",")
			ids = append(ids, strings.Split(trimmed, ",")...)
		}

		for _, id := range ids {
			id, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println("Failed to parse argument 🥺", id)
				os.Exit(1)
			}
	
			err = db.DeleteTask(tasks[id - 1].Key)
			if err != nil {
				fmt.Println("Failed to delete task", id, err)
				os.Exit(1)
			}
			fmt.Println("Completed task 🌸", id)
		}
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}
