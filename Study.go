package main

import "fmt"

func main() {
	pass := 7055
	if pass == 7055 {
		fmt.Scan(&pass)
		fmt.Println("Пароль верный")
	} else {
		fmt.Println("Пароль неверный")

	}
}
