// Package todo provides a simple todo list implementation.
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// Save writes the todo list to a file in JSON format
func (l *List) Save(filename string) error {
	data, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

// Load reads a todo list from a file
func (l *List) Load(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, l)
}

// Item represents a todo item with text and completion status.
type Item struct {
	Text string
	Done bool
}

// NewItem creates a new todo item with the specified text.
func NewItem(text string) Item {
	return Item{
		Text: text,
		Done: false,
	}
}

// List represents a collection of todo items.
type List struct {
	Items []Item
}

// NewList creates a new empty todo list.
func NewList() *List {
	return &List{
		Items: []Item{},
	}
}

// Add adds a new item to the list.
func (l *List) Add(text string) {
	item := NewItem(text)
	l.Items = append(l.Items, item)
}

// Complete marks the item at the specified index as done.
func (l *List) Complete(index int) error {
	if index < 0 || index >= len(l.Items) {
		return errors.New("item index out of range")
	}

	l.Items[index].Done = true
	return nil
}

// String returns a formatted string representation of the list.
func (l *List) String() string {
	if len(l.Items) == 0 {
		return "No items in the todo list"
	}

	result := "Todo List:\n"
	for i, item := range l.Items {
		status := " "
		if item.Done {
			status = "✓"
		}
		result += fmt.Sprintf("%d. [%s] %s\n", i+1, status, item.Text)
	}

	return result
}
