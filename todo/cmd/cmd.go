package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/TodoList/todo"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "todo [command]",
	Short: "A todo list application",
	Long:  `A Fast and Flexible todo list application built with Go`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println("Welcome to TodoList CLI Application!!")
	},
}

// AddCmd represents the add command
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add a new todo to the list`,
	Run: func(cmd *cobra.Command, args []string) {
		td := todo.Service()
		td.AddTodo(strings.Join(args, ""))
	},
}

// RemoveCmd represents the remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a todo",
	Long:  `Remove a todo from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		td := todo.Service()
		fmt.Printf("value :%v", args[0])
		if len(args[1:]) > 0 {
			fmt.Println("Invalid command")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid command")
			return
		}
		td.RemoveTodo(id)
	},
}
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lits a todo",
	Long:  `List all todos`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args[0:]) > 0 {
			fmt.Println("Invalid command")
			return
		}
		td := todo.Service()
		td.ListTodo()
	},
}

// Add the add command to the root command
func init() {
	RootCmd.AddCommand(AddCmd, RemoveCmd, ListCmd)
}
