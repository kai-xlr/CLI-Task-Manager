// Package todo provides a simple todo list implementation
package todo

// Item represents a todo item with text and completion status
type Item struct {
	Text string
	Done bool
}

// NewItem creates a new todo item with the specified text
func NewItem(text string) Item {
	return Item{
		Text: text,
		Done: false,
	}
}
