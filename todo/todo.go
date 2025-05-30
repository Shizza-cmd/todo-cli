package todo

import "fmt"

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

type TodoList []Todo

func (t *TodoList) Add(text string) {
	id := len(*t) + 1
	*t = append(*t, Todo{ID: id, Text: text, Done: false})
}

func (t *TodoList) List() []Todo {
	return *t
}

func (t *TodoList) MarkDone(id int) error {
	for i := range *t {
		if (*t)[i].ID == id {
			(*t)[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}

func (t *TodoList) Delete(id int) error {
	for i := range *t {
		if (*t)[i].ID == id {
			*t = append((*t)[:i], (*t)[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}
