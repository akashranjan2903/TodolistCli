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

func (t *todoList) AddTodo(str string) error {
	newItem := item{
		Task:       str,
		Complete:   true,
		UpodatedAt: time.Now(),
	}

	t.todoStore = append(t.todoStore, newItem)
	t.SavetoJson()
	return nil
}
func (t *todoList) RemoveTodo(str string) error {

	for key, v := range t.todoStore {

		if v.Task == str {
			t.todoStore = append(t.todoStore[:key], t.todoStore[key+1:]...)
		}

	}
	t.SavetoJson()
	return nil
}

func (t *todoList) ListTodo() []string {

	var list = []string{}
	for _, v := range t.todoStore {

		list = append(list, v.Task)

	}
	return list
}
