package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	a := []int{4, 8, 15, 16, 23, 42, 4, 15}
	b := []int{4, 11, 14, 8, 2, 2}
	searchA := make(map[int]bool)
	searchB := make(map[int]bool)
	for _, i := range a {
		searchA[i] = true
	}
	for _, i := range b {
		searchB[i] = true
	}
	var uniqA []int
	var uniqB []int
	for num := range searchA {
		uniqA = append(uniqA, num)
	}
	for num := range searchB {
		uniqB = append(uniqB, num)
	}
	var intert []int
	for k := range searchA {
		if searchB[k] {
			intert = append(intert, k)
		}
	}
	comma := strings.Trim(fmt.Sprint(uniqA), "[]") //check later
	comma = strings.ReplaceAll(comma, " ", ",")    // something wrong here
	fmt.Println(comma)
	sort.Ints(intert)
	sort.Ints(uniqA)
	sort.Ints(uniqB)
	fmt.Println(uniqA, ",", uniqB)
	fmt.Println(intert)

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
