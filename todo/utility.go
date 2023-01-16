package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

func (t todoList) SavetoJson() {
	// create files
	data, err := json.Marshal(t.todoStore)
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("db.json", data, 0o644); err != nil {
		panic(err)
	}
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
func (t todoList) AddnewId() int {
	if len(t.todoStore) == 0 {
		return 1
	}
	return t.todoStore[len(t.todoStore)-1].Id + 1
}
func (t *todoList) LoadFromJson() {
	// Check if the file exists
	if _, err := os.Stat("db.json"); os.IsNotExist(err) {
		os.Create("db.json")
	}

	// convert the file to a byte array
	data, err := os.ReadFile("db.json")
	if err != nil {
		panic(err)
	}

	if len(data) > 0 {
		// Unmarshal the data from the file to t.todoStore
		err = json.Unmarshal(data, &t.todoStore)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Println("No data found")
	}
}
