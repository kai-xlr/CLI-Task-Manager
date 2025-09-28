// Package todo provides a simple and efficient todo list implementation with
// persistent storage capabilities.
package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// Item represents a todo item with text, completion status, and metadata.
type Item struct {
	Text      string    `json:"text"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// NewItem creates a new todo item with the specified text.
// The item is created with the current timestamp and marked as not done.
func NewItem(text string) Item {
	text = strings.TrimSpace(text)
	if text == "" {
		text = "Untitled task"
	}
	
	return Item{
		Text:      text,
		Done:      false,
		CreatedAt: time.Now(),
	}
}

// Complete marks the item as done and sets the completion timestamp.
func (i *Item) Complete() {
	if !i.Done {
		i.Done = true
		now := time.Now()
		i.CompletedAt = &now
	}
}

// Uncomplete marks the item as not done and clears the completion timestamp.
func (i *Item) Uncomplete() {
	i.Done = false
	i.CompletedAt = nil
}

// String returns a formatted string representation of the item.
func (i Item) String() string {
	status := " "
	if i.Done {
		status = "✓"
	}
	return fmt.Sprintf("[%s] %s", status, i.Text)
}

// List represents a collection of todo items with management operations.
type List struct {
	Items []Item `json:"items"`
}

// NewList creates a new empty todo list.
func NewList() *List {
	return &List{
		Items: make([]Item, 0),
	}
}

// Add adds a new item to the list with the given text.
// Returns the index of the newly added item.
func (l *List) Add(text string) int {
	item := NewItem(text)
	l.Items = append(l.Items, item)
	return len(l.Items) - 1
}

// Complete marks the item at the specified index as done.
// Returns an error if the index is out of range.
func (l *List) Complete(index int) error {
	if err := l.validateIndex(index); err != nil {
		return err
	}

	l.Items[index].Complete()
	return nil
}

// Uncomplete marks the item at the specified index as not done.
// Returns an error if the index is out of range.
func (l *List) Uncomplete(index int) error {
	if err := l.validateIndex(index); err != nil {
		return err
	}

	l.Items[index].Uncomplete()
	return nil
}

// Delete removes the item at the specified index from the list.
// Returns an error if the index is out of range.
func (l *List) Delete(index int) error {
	if err := l.validateIndex(index); err != nil {
		return err
	}

	// Remove item by slicing around it
	l.Items = append(l.Items[:index], l.Items[index+1:]...)
	return nil
}

// Edit updates the text of the item at the specified index.
// Returns an error if the index is out of range.
func (l *List) Edit(index int, newText string) error {
	if err := l.validateIndex(index); err != nil {
		return err
	}

	newText = strings.TrimSpace(newText)
	if newText == "" {
		return errors.New("task text cannot be empty")
	}

	l.Items[index].Text = newText
	return nil
}

// Count returns the total number of items in the list.
func (l *List) Count() int {
	return len(l.Items)
}

// CountCompleted returns the number of completed items.
func (l *List) CountCompleted() int {
	count := 0
	for _, item := range l.Items {
		if item.Done {
			count++
		}
	}
	return count
}

// CountPending returns the number of pending (not completed) items.
func (l *List) CountPending() int {
	return l.Count() - l.CountCompleted()
}

// Clear removes all items from the list.
func (l *List) Clear() {
	l.Items = make([]Item, 0)
}

// validateIndex checks if the given index is valid for the current list.
func (l *List) validateIndex(index int) error {
	if index < 0 || index >= len(l.Items) {
		return fmt.Errorf("item index %d out of range (0-%d)", index, len(l.Items)-1)
	}
	return nil
}

// String returns a formatted string representation of the list.
func (l *List) String() string {
	if len(l.Items) == 0 {
		return "No items in the todo list\n"
	}

	completed := l.CountCompleted()
	total := l.Count()

	result := fmt.Sprintf("Todo List (%d/%d completed):\n", completed, total)
	for i, item := range l.Items {
		result += fmt.Sprintf("%d. %s\n", i+1, item.String())
	}

	return result
}

// Save writes the todo list to a file in JSON format with proper formatting.
// Returns an error if the file cannot be written or JSON marshaling fails.
func (l *List) Save(filename string) error {
	if filename == "" {
		return errors.New("filename cannot be empty")
	}

	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal todo list: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filename, err)
	}

	return nil
}

// Load reads a todo list from a JSON file.
// Returns an error if the file cannot be read or JSON unmarshaling fails.
func (l *List) Load(filename string) error {
	if filename == "" {
		return errors.New("filename cannot be empty")
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	if err := json.Unmarshal(data, l); err != nil {
		return fmt.Errorf("failed to unmarshal todo list from %s: %w", filename, err)
	}

	return nil
}
