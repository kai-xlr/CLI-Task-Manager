package todo

import (
	"os"
	"strings"
	"testing"
	"time"
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

func TestSaveAndLoad(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("", "todo-test")
	if err != nil {
		t.Fatalf("Could not create temp file: %v", err)
	}
	// Tear down
	defer os.Remove(tmpfile.Name())

	// Create and save a list
	list := NewList()
	list.Add("Task 1")
	list.Add("Task 2")
	list.Complete(0)

	if err := list.Save(tmpfile.Name()); err != nil {
		t.Fatalf("Failed to save list: %v", err)
	}

	// Load the list from the file
	loadedList := NewList()
	if err := loadedList.Load(tmpfile.Name()); err != nil {
		t.Fatalf("Failed to load list: %v", err)
	}

	// Verify the loaded list matches the original
	if len(loadedList.Items) != 2 {
		t.Errorf("Expected 2 items, got %d", len(loadedList.Items))
	}

	if !loadedList.Items[0].Done {
		t.Error("Expected first item to be completed")
	}

	if loadedList.Items[0].Text != "Task 1" {
		t.Errorf("Expected text 'Task 1', got '%s'", loadedList.Items[0].Text)
	}
}

// Test new Item methods

func TestNewItemWithTimestamp(t *testing.T) {
	start := time.Now()
	item := NewItem("Test task")
	end := time.Now()

	if item.Text != "Test task" {
		t.Errorf("Expected text 'Test task', got '%s'", item.Text)
	}

	if item.Done {
		t.Error("New item should not be marked as done")
	}

	if item.CreatedAt.Before(start) || item.CreatedAt.After(end) {
		t.Error("Created timestamp should be between test start and end")
	}

	if item.CompletedAt != nil {
		t.Error("New item should not have completion timestamp")
	}
}

func TestNewItemWithEmptyText(t *testing.T) {
	item := NewItem("")
	if item.Text != "Untitled task" {
		t.Errorf("Expected 'Untitled task' for empty text, got '%s'", item.Text)
	}

	item2 := NewItem("   ")
	if item2.Text != "Untitled task" {
		t.Errorf("Expected 'Untitled task' for whitespace text, got '%s'", item2.Text)
	}
}

func TestNewItemTrimsWhitespace(t *testing.T) {
	item := NewItem("  Test task  ")
	if item.Text != "Test task" {
		t.Errorf("Expected 'Test task', got '%s'", item.Text)
	}
}

func TestItemComplete(t *testing.T) {
	item := NewItem("Test task")
	
	// Initially not completed
	if item.Done {
		t.Error("Item should not be completed initially")
	}
	if item.CompletedAt != nil {
		t.Error("Item should not have completion timestamp initially")
	}

	// Complete the item
	start := time.Now()
	item.Complete()
	end := time.Now()

	if !item.Done {
		t.Error("Item should be marked as done after Complete()")
	}

	if item.CompletedAt == nil {
		t.Error("Item should have completion timestamp after Complete()")
	} else {
		if item.CompletedAt.Before(start) || item.CompletedAt.After(end) {
			t.Error("Completion timestamp should be between test start and end")
		}
	}

	// Complete again (should not change timestamp)
	originalTime := *item.CompletedAt
	time.Sleep(1 * time.Millisecond) // Ensure time difference
	item.Complete()

	if !item.CompletedAt.Equal(originalTime) {
		t.Error("Completion timestamp should not change on second Complete()")
	}
}

func TestItemUncomplete(t *testing.T) {
	item := NewItem("Test task")
	item.Complete()

	// Verify it's completed
	if !item.Done {
		t.Error("Item should be completed before uncomplete test")
	}
	if item.CompletedAt == nil {
		t.Error("Item should have completion timestamp before uncomplete test")
	}

	// Uncomplete the item
	item.Uncomplete()

	if item.Done {
		t.Error("Item should not be done after Uncomplete()")
	}
	if item.CompletedAt != nil {
		t.Error("Item should not have completion timestamp after Uncomplete()")
	}
}

func TestItemString(t *testing.T) {
	item := NewItem("Test task")
	expected := "[ ] Test task"
	if item.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, item.String())
	}

	item.Complete()
	expected = "[✓] Test task"
	if item.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, item.String())
	}
}

// Test new List methods

func TestListAddReturnsIndex(t *testing.T) {
	list := NewList()
	
	index1 := list.Add("First task")
	if index1 != 0 {
		t.Errorf("Expected index 0, got %d", index1)
	}

	index2 := list.Add("Second task")
	if index2 != 1 {
		t.Errorf("Expected index 1, got %d", index2)
	}
}

func TestListUncomplete(t *testing.T) {
	list := NewList()
	list.Add("Test task")
	list.Complete(0)

	// Verify it's completed
	if !list.Items[0].Done {
		t.Error("Task should be completed before uncomplete test")
	}

	// Uncomplete it
	err := list.Uncomplete(0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if list.Items[0].Done {
		t.Error("Task should not be completed after uncomplete")
	}

	// Test invalid index
	err = list.Uncomplete(1)
	if err == nil {
		t.Error("Expected error for invalid index")
	}
}

func TestListDelete(t *testing.T) {
	list := NewList()
	list.Add("Task 1")
	list.Add("Task 2")
	list.Add("Task 3")

	// Delete middle item
	err := list.Delete(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(list.Items) != 2 {
		t.Errorf("Expected 2 items after delete, got %d", len(list.Items))
	}

	if list.Items[0].Text != "Task 1" {
		t.Errorf("Expected 'Task 1', got '%s'", list.Items[0].Text)
	}

	if list.Items[1].Text != "Task 3" {
		t.Errorf("Expected 'Task 3', got '%s'", list.Items[1].Text)
	}

	// Test invalid index
	err = list.Delete(5)
	if err == nil {
		t.Error("Expected error for invalid index")
	}
}

func TestListEdit(t *testing.T) {
	list := NewList()
	list.Add("Original task")

	// Edit the task
	err := list.Edit(0, "Updated task")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if list.Items[0].Text != "Updated task" {
		t.Errorf("Expected 'Updated task', got '%s'", list.Items[0].Text)
	}

	// Test with whitespace
	err = list.Edit(0, "  Trimmed task  ")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if list.Items[0].Text != "Trimmed task" {
		t.Errorf("Expected 'Trimmed task', got '%s'", list.Items[0].Text)
	}

	// Test with empty text
	err = list.Edit(0, "")
	if err == nil {
		t.Error("Expected error for empty text")
	}

	err = list.Edit(0, "   ")
	if err == nil {
		t.Error("Expected error for whitespace-only text")
	}

	// Test invalid index
	err = list.Edit(1, "New text")
	if err == nil {
		t.Error("Expected error for invalid index")
	}
}

func TestListCount(t *testing.T) {
	list := NewList()
	if list.Count() != 0 {
		t.Errorf("Expected count 0, got %d", list.Count())
	}

	list.Add("Task 1")
	if list.Count() != 1 {
		t.Errorf("Expected count 1, got %d", list.Count())
	}

	list.Add("Task 2")
	if list.Count() != 2 {
		t.Errorf("Expected count 2, got %d", list.Count())
	}
}

func TestListCountCompleted(t *testing.T) {
	list := NewList()
	if list.CountCompleted() != 0 {
		t.Errorf("Expected completed count 0, got %d", list.CountCompleted())
	}

	list.Add("Task 1")
	list.Add("Task 2")
	list.Add("Task 3")

	if list.CountCompleted() != 0 {
		t.Errorf("Expected completed count 0, got %d", list.CountCompleted())
	}

	list.Complete(0)
	if list.CountCompleted() != 1 {
		t.Errorf("Expected completed count 1, got %d", list.CountCompleted())
	}

	list.Complete(2)
	if list.CountCompleted() != 2 {
		t.Errorf("Expected completed count 2, got %d", list.CountCompleted())
	}
}

func TestListCountPending(t *testing.T) {
	list := NewList()
	if list.CountPending() != 0 {
		t.Errorf("Expected pending count 0, got %d", list.CountPending())
	}

	list.Add("Task 1")
	list.Add("Task 2")
	list.Add("Task 3")

	if list.CountPending() != 3 {
		t.Errorf("Expected pending count 3, got %d", list.CountPending())
	}

	list.Complete(0)
	if list.CountPending() != 2 {
		t.Errorf("Expected pending count 2, got %d", list.CountPending())
	}

	list.Complete(1)
	list.Complete(2)
	if list.CountPending() != 0 {
		t.Errorf("Expected pending count 0, got %d", list.CountPending())
	}
}

func TestListClear(t *testing.T) {
	list := NewList()
	list.Add("Task 1")
	list.Add("Task 2")
	list.Add("Task 3")

	if list.Count() != 3 {
		t.Errorf("Expected count 3 before clear, got %d", list.Count())
	}

	list.Clear()

	if list.Count() != 0 {
		t.Errorf("Expected count 0 after clear, got %d", list.Count())
	}

	if len(list.Items) != 0 {
		t.Errorf("Expected items slice to be empty after clear, got length %d", len(list.Items))
	}
}

func TestListStringFormatting(t *testing.T) {
	// Test empty list
	list := NewList()
	result := list.String()
	expected := "No items in the todo list\n"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	// Test list with items
	list.Add("Task 1")
	list.Add("Task 2")
	list.Complete(0)

	result = list.String()
	if !strings.Contains(result, "Todo List (1/2 completed):") {
		t.Error("String output should contain completion statistics")
	}

	if !strings.Contains(result, "1. [✓] Task 1") {
		t.Error("String output should contain completed task with checkmark")
	}

	if !strings.Contains(result, "2. [ ] Task 2") {
		t.Error("String output should contain pending task with empty checkbox")
	}
}

// Test persistence error handling

func TestSaveWithEmptyFilename(t *testing.T) {
	list := NewList()
	err := list.Save("")
	if err == nil {
		t.Error("Expected error when saving with empty filename")
	}
}

func TestLoadWithEmptyFilename(t *testing.T) {
	list := NewList()
	err := list.Load("")
	if err == nil {
		t.Error("Expected error when loading with empty filename")
	}
}

func TestLoadNonexistentFile(t *testing.T) {
	list := NewList()
	err := list.Load("/nonexistent/path/file.json")
	if err == nil {
		t.Error("Expected error when loading nonexistent file")
	}
}
