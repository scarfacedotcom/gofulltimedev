package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	result, err := thirdpartyHTTPCall()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the response took %v -> %+v\n", time.Since(start), result)

}

func fetchUserID() (string, error) {
	ctx, cancel := context.WithTimeout(time.Millisecond * 100)
	defer cancel()

	type result struct {
		userID string
		err    error
	}

	resultch := make(chan result, 1)

	go func() {
		res, err := thirdpartyHTTPCall()
		resultch <- result{
			userID: res,
			err:    err,
		}
	}()
}

func thirdpartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 100)
	return "user id 1", nil
}
