package async_worker

import (
	"context"
	"fmt"
	"time"
)

// TaskStatus defines a struct for task status management.
type TaskStatus struct {
	id        string // Unique identifier for the task
	status    string
	createdAt time.Time // Timestamp when the task was created
	// Add more common attributes as needed
}

// NewTaskStatus creates a new TaskStatus instance with default values.
func NewTaskStatus() *TaskStatus {
	return &TaskStatus{
		id:        GenerateTaskID(), // Generate a unique ID for the task
		createdAt: time.Now(),       // Set the creation timestamp
	}
}

// PreProcess executes common pre-processing steps for all tasks.
func (ts *TaskStatus) PreProcess(ctx context.Context) error {
	// Example: Logging before task execution
	ts.status = "pending"
	fmt.Printf("[%s] Starting task execution...\n Status: %s\n", ts.id, ts.status)

	// Example: Adding common context values
	ctx = context.WithValue(ctx, "TaskID", ts.id)
	ctx = context.WithValue(ctx, "TaskCreatedAt", ts.createdAt)

	// Example: Additional common operations...

	return nil
}

// PostProcess executes common post-processing steps for all tasks.
func (ts *TaskStatus) PostProcess(ctx context.Context) error {
	ts.status = "completed"
	fmt.Printf("[%s] Task execution completed.\n Status: %s\n", ts.id, ts.status)

	// Example: Additional post-processing steps...

	return nil
}

// SetErrorStatus sets an error status and logs the error message.
func (ts *TaskStatus) SetErrorStatus(err error) {
	ts.status = "failed"
	fmt.Printf("[%s] Task execution failed: %s\n", ts.id, err.Error())
}

// generateTaskID generates a unique ID for each task.
func GenerateTaskID() string {
	return fmt.Sprintf("Task-%d", time.Now().UnixNano())
}
