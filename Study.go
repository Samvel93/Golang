package main

import "fmt"

func main() {
	a := [4]int{5, 1, 2, 5}
	b := [6]int{4, 2, 5, 1, 1, 2}
	var x []int = a[0:3]
	var y []int = b[0:4]
	var z []int = append(x, b[0])
	fmt.Println(x, ",", y)
	fmt.Println(x)
	fmt.Println(z)

}

/*Входные данные: a = [5, 1, 2, 5], b = [4, 2, 5, 1, 1, 2]
Выходные данные:
[5, 1, 2], [4, 2, 5, 1]
[5, 1, 2]
[5, 1, 2, 4]
*/
