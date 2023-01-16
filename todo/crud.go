package todo

import "time"

type todotype interface {
	AddTodo(str string) error
	RemoveTodo(str string) error
	ListTodo() []string
}

type item struct {
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
	newItem := item{
		Task:       str,
		Complete:   true,
		UpodatedAt: time.Now(),
	}

	td.todoStore = append(td.todoStore, newItem)
	td.SavetoJson()
	return nil
}
func (td *todoList) RemoveTodo(str string) error {

	for key, v := range td.todoStore {

		if v.Task == str {
			td.todoStore = append(td.todoStore[:key], td.todoStore[key+1:]...)
		}

	}
	td.SavetoJson()
	return nil
}

func (td *todoList) ListTodo() []string {

	var list = []string{}
	for _, v := range td.todoStore {

		list = append(list, v.Task)

	}
	return list
}
