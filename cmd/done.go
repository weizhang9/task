package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"gophercises/task/db"

	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Makrs task(s) as done. Use comma or space to separate multiple tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("You have no tasks to delete 🥶")
			os.Exit(1)
		}

		if isFlagPassed("all") || isFlagPassed("a") {
			err := db.DeleteAllTasks(tasks)
			if err != nil {
				fmt.Println("Failed to complete all tasks", err)
				os.Exit(1)
			}
			fmt.Println("Completed all tasks ヾ(✿◔ ڼ ◔ )ノ🌸")
			os.Exit(0)
		}

		var ids []string
		for _, arg := range args {
			trimmed := strings.Trim(arg, ",")
			ids = append(ids, strings.Split(trimmed, ",")...)
		}

		for _, id := range ids {
			trimmed := strings.TrimSpace(id)
			id, err := strconv.Atoi(trimmed)
			if err != nil {
				fmt.Println("Failed to parse argument 🥺", trimmed)
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
	flag.BoolP("all", "a", false, "Marks all task(s) as done")
	flag.Parse()
	RootCmd.AddCommand(doneCmd)
}

func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}