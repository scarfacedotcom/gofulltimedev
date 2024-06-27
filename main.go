package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.Background()
	userID := 19
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("value: ", val)
	fmt.Println("took us", time.Since(start))

}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	value, err := fetchThirdPartyStuffWhichCanBelow()
	if err != nil {
		return 0, err
	}
	return value, nil
}

func fetchThirdPartyStuffWhichCanBelow() (int, error) {
	time.Sleep(2 * time.Microsecond)

	return 666, nil
}
