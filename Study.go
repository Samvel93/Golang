package main

import "fmt"

func main() {
	var result int
	fmt.Scan(&result)
	if result >= 3 {
		fmt.Println("Победа")
	} else if result == 3 {
		fmt.Println("Ничья")
	} else {
		fmt.Println("Поражение")
	}

}

/*Входные данные: a = [5, 1, 2, 5], b = [4, 2, 5, 1, 1, 2]
Выходные данные:
[5, 1, 2], [4, 2, 5, 1]
[5, 1, 2]
[5, 1, 2, 4]
*/
