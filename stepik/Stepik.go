package main

import "fmt"

func main() {
	numbers := make([]int, 5)
	for i := 0; i < cap(numbers); i++ {
		numbers[i] = i
	}
	fmt.Println(numbers)
}

/*list := []string{"apple", "banana", "cherry", "date", "elderberry"}
	list[2] = "citrus"
	controlist := list[:]
	list2 := []string{"orange", "papaya"}
	fmt.Printf("%q\n", append(controlist, list2...))
	fmt.Printf("%q\n", append(append(controlist[:1], controlist[2:]...), list2...))

/*num1 := 1
	num2 := 10
	for num1 <= num2 {

		fmt.Print(num2, "...")
		num2--
		if num1 == 0 || num1 == 11 {
			break
		}
	}
	fmt.Println("Пуск")
 {
	var age int
	fmt.Scan(&age)
	switch {
	case age >= 75:
		fmt.Println("Very old")
	case age < 75 && age >= 60:
		fmt.Println("Old")
	case age < 60 && age >= 39:
		fmt.Println("Middle aged")
	case age < 39 && age >= 18:
		fmt.Println("Young")
	default:
		fmt.Println("Child")
	}
*/
