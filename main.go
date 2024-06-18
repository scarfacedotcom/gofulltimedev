package main

import (
	"fmt"
	"time"
)

func fetchResource(n int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("result %d", n)
}

func main() {

	resultch := make(chan string)

	go func() {
		result := <-resultch
		fmt.Println(result)
	}()
	resultch <- "foo"

}
