package todo

import (
	"encoding/json"
	"os"
)

const DataFile = "data/todos.json"

func Load() (TodoList, error) {
	data, err := os.ReadFile(DataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return TodoList{}, nil
		}

		return nil, err
	}

	var todos TodoList
	err = json.Unmarshal(data, &todos)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t *TodoList) Save() error {
	if _, err := os.Stat("data"); os.IsNotExist(err) {
		err := os.Mkdir("data", 0755)
		if err != nil {
			return err
		}
	}

	data, _ := json.MarshalIndent(t, "", "  ")
	return os.WriteFile(DataFile, data, 0644)
}
