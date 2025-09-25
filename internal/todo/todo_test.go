package todo

import (
	"testing"
)

func TestNewItem(t *testing.T) {
	item := NewItem("Learn Go")

	if item.Text != "Learn Go" {
		t.Errorf("Expected text to be 'Learn Go', got '%s'", item.Text)
	}

	if item.Done {
		t.Error("New item should not be marked as done")
	}
}
