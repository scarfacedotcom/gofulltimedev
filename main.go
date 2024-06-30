package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "foo", "boo")
	userID := 19
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value: ", val)
	fmt.Println("took us", time.Since(start))

}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	val := ctx.Value("foo")
	fmt.Println(val.(string))
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()
	respchn := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBelow()
		respchn <- Response{
			value: val,
			err:   err,
		}

	}()

	for {
		select {
		case <-ctx.Done():
			return 0,
				fmt.Errorf("took too long to print")
		case resp := <-respchn:
			return resp.value, resp.err
		}
	}

}

func fetchThirdPartyStuffWhichCanBelow() (int, error) {
	time.Sleep(2 * time.Microsecond)

	return 666, nil
}
