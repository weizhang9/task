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
	"os"

	"github.com/boltdb/bolt"
	"github.com/spf13/cobra"
)

type database struct {
	name string
	port os.FileMode
	db   *bolt.DB
	err  error
}

var taskDB = database{
	name: "todos.db",
	port: 0600,
}

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A CLI to make your todo list",
	Long: `A CLI tool that can be used to manage your TODOs in the terminal. 
	
	You can use this tool to add a new TODO;
	list out all the incomplete TODOs;
	or mark a TODO as complete.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkErr(e error, info string) {
	if e != nil {
		log.Fatalln(info, e)
	}
}

func connectDB() {
	taskDB.db, taskDB.err = bolt.Open(taskDB.name, taskDB.port, nil)
	checkErr(taskDB.err, "[Fail to connect to DB]")
}