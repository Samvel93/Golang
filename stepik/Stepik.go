package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	if x > 1000 {
		fmt.Println("Apple")
	} else if x >= 500 && x <= 1000 {
		fmt.Println("Samsung")
	} else {
		fmt.Println("Nokia")
	}
}
