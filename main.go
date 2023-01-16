package main

import (
	"fmt"
	"os"

	"github.com/TodoList/todo/cmd"
)

func main() {

	// Normal execution

	// newservice := todo.Service()

	// add

	// newservice.AddTodo("temp")

	// newservice.AddTodo("test")

	// newservice.RemoveTodo(1)
	// fmt.Println(newservice.ListTodo())

	// using command line

	// args := os.Args[1:]
	// switch args[0] {
	// case "add":
	// 	fmt.Println("Add todo")
	// 	newservice.AddTodo(strings.Join(args[1:], ""))
	// case "rm":
	// 	fmt.Println("Remove todo")
	// 	id, err := strconv.Atoi(strings.Join(args[1:], ""))
	// 	if err != nil {
	// 		fmt.Println("Invalid command")
	// 		return
	// 	}
	// 	newservice.RemoveTodo(id)
	// case "list":
	// 	if len(args[1:]) > 0 {
	// 		fmt.Println("Invalid command")
	// 		return
	// 	}
	// 	newservice.ListTodo()
	// default:
	// 	fmt.Println("Invalid command")
	// }

	// using cobra cli

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
