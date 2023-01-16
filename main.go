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
	fmt.Println(newservice.AddTodo("Akash"))

	fmt.Println(newservice.RemoveTodo("Akash"))
	fmt.Println(newservice.ListTodo())

}
