package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func (t todoList) SavetoJson() {
	// create file
	filename := "db.json"
	err := checkFile(filename)
	if err != nil {
		panic(err)
	}
	file, err := os.ReadFile("db.json")
	if err != nil {
		panic(err)
	}

	prevdata := []item{}
	json.Unmarshal(file, &prevdata)

	fmt.Printf("prvious value%v", prevdata)
	newdata := append(prevdata, t.todoStore...)

	//  marashal the data from  new data
	newdatabyte, err := json.Marshal(newdata)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, newdatabyte, 0644)
	if err != nil {
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

func (t todoList) LoadFromJson() {
	// Check if the file exists
	if _, err := os.Stat("db.json"); os.IsNotExist(err) {
		os.Create("db.json")
	}

	// Open the file
	file, err := os.OpenFile("db.json", os.O_RDONLY, 0o644)
	if err != nil {
		panic(err)
	}

	// convert the file to a byte array
	fileByte, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	if len(fileByte) > 0 {
		// Unmarshal the data from the file to t.todoStore
		err = json.Unmarshal([]byte(fileByte), &t.todoStore)
		if err != nil {
			panic(err)
		}
	}

	// Close the file
	err = file.Close()
	if err != nil {
		panic(err)
	}
}
