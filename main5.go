package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// input: [2,3,4], n=> n *2
//Output [4,6,8]

func MappedValues[T constraints.Ordered](values []T, mapFunc func(T) T) []T {
	var newValues []T
	for _, v := range values {
		newValue := mapFunc(v)
		newValues = append(newValues, newValue)
	}
	return newValues

}

func main5() {

	result := MappedValues([]int{1, 2, 3, 4}, func(n int) int {
		return n * 5
	})

	fmt.Println("result:", result)

}
