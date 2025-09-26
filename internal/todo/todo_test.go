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

func TestAddItem(t *testing.T) {
	list := NewList()
	list.Add("Buy Milk")

	if len(list.Items) != 1 {
		t.Errorf("Expected list to have 1 item, got %d", len(list.Items))
	}

	if list.Items[0].Text != "Buy Milk" {
		t.Errorf("Expected item text to be 'Buy Milk', got '%s'", list.Items[0].Text)
	}
}

func TestCompleteItem(t *testing.T) {
	list := NewList()
	list.Add("Write Code")

	err := list.Complete(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !list.Items[0].Done {
		t.Error("Item should be marked as done")
	}

	err = list.Complete(1)
	if err == nil {
		t.Errorf("Expected error when completing non-existent item")
	}
}
