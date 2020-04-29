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
	"strconv"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [\"a TODO string\"]",
	Short: "Add a new TODO",
	Long: `add command allows you to add new TODO(s) to the TODO list; each TODO must be in double quotes. For example:

	$ task add "write shopping list" "call mum"
	the above command would add "write shopping list" and "call mum" to your TODO list`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := addTodo(args)
		if err != nil {
			log.Fatalln("[Fail to add todo]", err)
		} else {
			fmt.Println("added successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func addTodo(todos []string) error {
	taskDB.db, taskDB.err = bolt.Open(taskDB.name, taskDB.port, nil)
	checkErr(taskDB.err, "[Fail to connect to DB]")
	defer taskDB.db.Close()

	if err := taskDB.db.Update(func(tx *bolt.Tx) error {
		// Create a bucket.
		b, err := tx.CreateBucketIfNotExists([]byte("todos"))
		if err != nil {
			return err
		}

		for _, v := range todos {
			lastk, _ := b.Cursor().Last()
			if lastk != nil {
				lastkInt, _ := strconv.Atoi(string(lastk))
				if err := b.Put([]byte(fmt.Sprint(lastkInt+1)), []byte(v)); err != nil {
					return err
				}
			} else {
				if err := b.Put([]byte(fmt.Sprint("1")), []byte(v)); err != nil {
					return err
				}
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
