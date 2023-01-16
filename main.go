package main

import (
	"fmt"

	"github.com/TodoList/todo"
)

func main() {

	newservice := todo.Service()
	// newser:

	// add

	fmt.Println(newservice.AddTodo("temp"))
	fmt.Println(newservice.AddTodo("test"))

	fmt.Println(newservice.RemoveTodo("test"))
	fmt.Println(newservice.ListTodo())

}
