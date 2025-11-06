package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{4, 8, 15, 16, 23, 42, 4, 15}
	searchA := make(map[int]bool)
	for _, i := range a {
		searchA[i] = true
	}
	var uniqA []int
	for num := range searchA {
		uniqA = append(uniqA, num)
	}
	b := []int{4, 11, 14, 8, 2, 2}
	searchB := make(map[int]bool)
	for _, i := range b {
		searchB[i] = true
	}
	var uniqB []int
	for num := range searchB {
		uniqB = append(uniqB, num)
	}
	sort.Ints(uniqA)
	sort.Ints(uniqB)
	fmt.Println(uniqA, uniqB)

}

/*
-Вернуть два слайса с уникальными элементами.
-Найти пересечения значений в двух слайсах.
-Вернуть слайс с уникальными элементами из слайсов a и b
Входные данные: a = [5, 1, 2, 5], b = [4, 2, 5, 1, 1, 2]
Выходные данные:
[5, 1, 2], [4, 2, 5, 1]
[5, 1, 2]
[5, 1, 2, 4]
*/
