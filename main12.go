package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main12() {
	start := time.Now()

	userID, err := fetchUserID()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("the response took %v -> %+v\n", time.Since(start), userID)

}

func fetchUserID() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
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

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userID, res.err
	}
}

func thirdpartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 300)
	return "user id 1", nil
}
