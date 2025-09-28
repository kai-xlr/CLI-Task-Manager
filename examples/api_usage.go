package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kai-xlr/CLI-Task-Manager/internal/todo"
)

// This example demonstrates how to use the todo package as a library
// in your own Go applications.
func main() {
	fmt.Println("=== Todo Package API Usage Example ===")
	fmt.Println()

	// Create a new todo list
	list := todo.NewList()
	fmt.Println("Created new todo list")

	// Add some tasks
	fmt.Println("\n1. Adding tasks:")
	tasks := []string{
		"Set up development environment",
		"Write unit tests",
		"Implement core features",
		"Add documentation",
		"Deploy to production",
	}

	for _, task := range tasks {
		index := list.Add(task)
		fmt.Printf("   Added: %s (index %d)\n", task, index)
	}

	// Display the list
	fmt.Println("\n2. Current todo list:")
	fmt.Print(list)

	// Complete some tasks
	fmt.Println("3. Completing some tasks:")
	if err := list.Complete(0); err != nil {
		log.Printf("Error completing task 0: %v", err)
	} else {
		fmt.Println("   ✓ Completed task 1")
	}

	if err := list.Complete(2); err != nil {
		log.Printf("Error completing task 2: %v", err)
	} else {
		fmt.Println("   ✓ Completed task 3")
	}

	// Edit a task
	fmt.Println("\n4. Editing a task:")
	if err := list.Edit(1, "Write comprehensive unit tests"); err != nil {
		log.Printf("Error editing task: %v", err)
	} else {
		fmt.Println("   ✓ Updated task 2")
	}

	// Show statistics
	fmt.Println("\n5. Statistics:")
	fmt.Printf("   Total tasks: %d\n", list.Count())
	fmt.Printf("   Completed: %d\n", list.CountCompleted())
	fmt.Printf("   Pending: %d\n", list.CountPending())

	// Display updated list
	fmt.Println("\n6. Updated todo list:")
	fmt.Print(list)

	// Save to file
	fmt.Println("7. Saving to file:")
	filename := "example_todos.json"
	if err := list.Save(filename); err != nil {
		log.Fatalf("Failed to save: %v", err)
	}
	fmt.Printf("   ✓ Saved to %s\n", filename)

	// Load from file (demonstrate loading)
	fmt.Println("\n8. Loading from file (demonstration):")
	newList := todo.NewList()
	if err := newList.Load(filename); err != nil {
		log.Fatalf("Failed to load: %v", err)
	}
	fmt.Printf("   ✓ Loaded from %s\n", filename)
	fmt.Print("   Loaded list:\n")
	fmt.Print(newList)

	// Demonstrate uncomplete functionality
	fmt.Println("9. Uncompleting a task:")
	if err := newList.Uncomplete(0); err != nil {
		log.Printf("Error uncompleting task: %v", err)
	} else {
		fmt.Println("   ✓ Marked task 1 as incomplete")
	}

	// Delete a task
	fmt.Println("\n10. Deleting a task:")
	if err := newList.Delete(4); err != nil {
		log.Printf("Error deleting task: %v", err)
	} else {
		fmt.Println("   ✓ Deleted task 5")
	}

	// Final state
	fmt.Println("\n11. Final state:")
	fmt.Print(newList)

	// Clean up example file
	if err := os.Remove(filename); err != nil {
		log.Printf("Warning: Could not remove example file: %v", err)
	} else {
		fmt.Printf("   ✓ Cleaned up %s\n", filename)
	}

	fmt.Println("\n=== Example completed! ===")
	fmt.Println("This demonstrates the core functionality of the todo package.")
	fmt.Println("You can integrate this package into your own Go applications.")
}