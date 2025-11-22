package main

import (
	"fmt"
)

func eraseindex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	s = eraseindex(s, 2)
	fmt.Println(s)
}
