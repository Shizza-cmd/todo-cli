package main

import (
	"flag"
	"fmt"
	"os"

	"shizza-cmd/todo-cli/todo"
)

func main() {
	add := flag.String("add", "", "Add new todo")
	list := flag.Bool("list", false, "List all todos")
	done := flag.Int("done", 0, "Mark todo as done")
	delete := flag.Int("delete", 0, "Delete a todo")

	flag.Parse()

	tl, err := todo.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading todos: %v\n", err)
		os.Exit(1)
	}

	switch {
	case *add != "":
		tl.Add(*add)
		tl.Save()
		fmt.Println("Added:", *add)
	case *list:
		for _, t := range tl.List() {
			status := " "
			if t.Done {
				status = "x"
			}
			fmt.Printf("[%s] %d. %s\n", status, t.ID, t.Text)
		}
	case *done > 0:
		if err := tl.MarkDone(*done); err != nil {
			fmt.Fprintln(os.Stderr, "Error marking done:", err)
		} else {
			tl.Save()
			fmt.Println("Marked todo", *done, "as done")
		}
	case *delete > 0:
		if err := tl.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, "Error deleting todo:", err)
		} else {
			tl.Save()
			fmt.Println("Deleted todo", *delete)
		}
	default:
		fmt.Println("No command provided.")
		fmt.Println("Usage:")
		fmt.Println("  -add \"text\"       Add new todo")
		fmt.Println("  -list             List all todos")
		fmt.Println("  -done ID          Mark todo as done")
		fmt.Println("  -delete ID        Delete a todo")
	}
}
