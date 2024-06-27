package main

import (
	"context"
	"fmt"
	"time"
)

func main13() {
	// Create a background context
	ctx := context.Background()

	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// Simulate a goroutine that performs some work
	go func() {
		// Simulate work by sleeping for 1 second
		time.Sleep(1 * time.Second)

		// Check if the context is still active
		select {
		case <-ctx.Done():
			// Context was canceled or timed out
			fmt.Println("Operation canceled or timed out:", ctx.Err())
			return
		default:
			// Perform the work
			fmt.Println("Work completed")
		}
	}()

	// Wait for the work to be done
	select {
	case <-ctx.Done():
		fmt.Println("Main function:", ctx.Err())
	}
}
