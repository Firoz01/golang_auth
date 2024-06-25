// worker_pool.go
package async_worker

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Task is an interface that any task should implement to be processed by the worker pool.
type Task interface {
	Process(ctx context.Context) error
}

// WorkerPool is a generic worker pool for processing tasks asynchronously.
type WorkerPool struct {
	taskChan chan Task
	wg       sync.WaitGroup
}

// NewWorkerPool creates a new WorkerPool with the given number of workers and buffer size.
func NewWorkerPool(bufferSize int) *WorkerPool {
	return &WorkerPool{
		taskChan: make(chan Task, bufferSize),
	}
}

// Start initializes the worker pool and starts the workers.
func (wp *WorkerPool) Start(ctx context.Context, numWorkers int) {
	wp.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wp.wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Worker stopped")
					return
				case task := <-wp.taskChan:
					if err := task.Process(ctx); err != nil {
						fmt.Println("Failed to process task:", err)
					}
				}
			}
		}()
	}
}

// Stop gracefully shuts down the worker pool.
func (wp *WorkerPool) Stop() error {
	// Close the taskChan to signal termination
	close(wp.taskChan)

	// Use a channel to receive errors
	errCh := make(chan error, 1)

	// Use a Goroutine to wait for the worker group to finish
	go func() {
		defer close(errCh)
		wp.wg.Wait()
	}()

	// Check if there's any error reported
	select {
	case <-time.After(time.Second): // Timeout after a certain period (adjust as needed)
		return fmt.Errorf("timeout waiting for worker to stop")
	case err := <-errCh:
		return err // Return the error received
	}
}

// Enqueue adds a task to the worker pool.
func (wp *WorkerPool) Enqueue(task Task) error {
	defer func() {
		if r := recover(); r != nil {
			// Handle the panic caused by closed channel
			fmt.Println("Channel closed. Failed to enqueue task:", r)
		}
	}()

	// Attempt to send to taskChan
	select {
	case wp.taskChan <- task:
		return nil
	default:
		// If taskChan is full, attempt to handle this error case
		return fmt.Errorf("failed to enqueue task: taskChan is full")
	}
}
