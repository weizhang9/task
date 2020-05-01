/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [TODO number 1, TODO number 2]",
	Short: "Mark an existing TODO as completed",
	Long: `done command allows you to mark existing TODO(s) as completed; each TODO number is separated by a comma. For example:

	if you have a list of TODO like:
	1. wash dishes
	2. take bins out
	you could:
	$ task done 1, 2
	which will mark TODO 1 and 2 as complete and remove it from the list`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, v := range args {
			if err := doneTodo(v); err != nil {
				log.Fatalf("TODO [%s] doesn't exist", v)
			} else {
				fmt.Println("Completed TODO:", v)
			}
		}
		
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}

func doneTodo(todo string)  error {
	connectDB()
	defer taskDB.db.Close()
	return done(todo)
}

func done(todo string) error {
	if err := taskDB.db.View(func(tx *bolt.Tx) error {
		v := tx.Bucket([]byte("todos")).Get([]byte(todo))
		// check if key exist
		if v == nil {
			return fmt.Errorf("TODO [%s] doesn't exist", todo)
		}

		// if exists, delete it
		return remove(todo)
	}); err != nil {
		return err
	}
	
	return nil
}

func remove(todo string) error {
	if err := taskDB.db.Update(func(tx *bolt.Tx) error {
		return tx.Bucket([]byte("todos")).Delete([]byte(todo))
	}); err != nil {
		return err
	}
	return nil
}