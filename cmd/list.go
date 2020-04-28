package cmd

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the incomplete TODOs",
	Long: `list command allows you to list out all the incomplete TODOs, it doesn't accept any arg. For example:

	$ task list
and the result would look like:
	1. wash dishes
	2. take bins out`,
	Run: func(cmd *cobra.Command, args []string) {
		checkErr(listTodo(), "Fail to list todos")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listTodo() error {
	taskDB.db, taskDB.err = bolt.Open(taskDB.name, taskDB.port, nil)
	checkErr(taskDB.err, "[Fail to connect to DB]")
	defer taskDB.db.Close()

	if err := taskDB.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("todos"))
		b.ForEach(func(k, v []byte) error {
			fmt.Printf("TODO [%s]: %s\n", string(k), string(v))
			return nil
		})
		return nil
	}); err != nil {
		return err
	}
	return nil
}