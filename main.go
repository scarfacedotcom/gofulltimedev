package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println(transformString("scar-face", Uppercase))
}
