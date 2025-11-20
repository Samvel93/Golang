package main

import "fmt"

func main() {
	old := []string{"one", "two", "three", "four", "five"}
	new := make([]string, len(old))
	copy(new, old)
	new = append(new, "six")
	fmt.Println(len(new))
	fmt.Println(len(old))
	fmt.Println(old)
	fmt.Println(new)
}
