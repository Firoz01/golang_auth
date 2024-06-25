package async_worker

import (
	"context"
	"fmt"
	"github.com/go-frame/internals/entity"
	"sync"
	"time"
)

type AsyncLogger interface {
	LogLoginAttempt(ctx context.Context, attempt *entity.LoginAttempt) error
}

type LoginAttemptWorker struct {
	logChan chan entity.LoginAttempt
	logger  AsyncLogger
	wg      sync.WaitGroup
}

func NewLoginAttemptWorker(logger AsyncLogger, bufferSize int) *LoginAttemptWorker {
	return &LoginAttemptWorker{
		logChan: make(chan entity.LoginAttempt, bufferSize),
		logger:  logger,
	}
}

func (w *LoginAttemptWorker) Start(ctx context.Context) {
	w.wg.Add(1)
	go func() {
		defer w.wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Login attempt worker stopped")
				return
			case attempt := <-w.logChan:
				if err := w.logger.LogLoginAttempt(ctx, &attempt); err != nil {
					fmt.Println("Failed to log login attempt:", err)
				}
			}
		}
	}()
}

func (w *LoginAttemptWorker) Stop() error {
	// Close the logChan to signal termination
	close(w.logChan)

	// Use a channel to receive errors
	errCh := make(chan error, 1)

	// Use a Goroutine to wait for the worker group to finish
	go func() {
		defer close(errCh)
		w.wg.Wait()
	}()

	// Check if there's any error reported
	select {
	case <-time.After(time.Second): // Timeout after a certain period (adjust as needed)
		return fmt.Errorf("timeout waiting for worker to stop")
	case err := <-errCh:
		return err // Return the error received
	}
}
func (w *LoginAttemptWorker) Enqueue(attempt entity.LoginAttempt) error {
	defer func() {
		if r := recover(); r != nil {
			// Handle the panic caused by closed channel
			fmt.Println("Channel closed. Failed to enqueue attempt:", r)
		}
	}()

	// Attempt to send to logChan
	select {
	case w.logChan <- attempt:
		return nil
	default:
		// If logChan is full, attempt to handle this error case
		return fmt.Errorf("failed to enqueue attempt: logChan is full")
	}
}
