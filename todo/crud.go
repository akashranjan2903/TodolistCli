package todo

import (
	"fmt"
	"time"
)

type todotype interface {
	AddTodo(str string) error
	RemoveTodo(id int) error
	ListTodo()
}

type item struct {
	Id         int
	Task       string
	Complete   bool
	UpodatedAt time.Time
}
type todoList struct {
	todoStore []item
}

func Service() todotype {
	return &todoList{}
}

func (td *todoList) AddTodo(str string) error {
	td.LoadFromJson()
	newItem := item{
		Id:         td.AddnewId(),
		Task:       str,
		Complete:   true,
		UpodatedAt: time.Now(),
	}

	td.todoStore = append(td.todoStore, newItem)

	td.SavetoJson()
	return nil
}
func (td *todoList) RemoveTodo(id int) error {
	td.LoadFromJson()
	for key, v := range td.todoStore {

		if v.Id == id {
			td.todoStore = append(td.todoStore[:key], td.todoStore[key+1:]...)
		}

	}
	td.SavetoJson()
	return nil
}

func (td *todoList) ListTodo() {
	td.LoadFromJson()
	var list = []string{}
	for _, v := range td.todoStore {

		list = append(list, v.Task)

	}
	fmt.Printf("All list todos :%v", list)
}
