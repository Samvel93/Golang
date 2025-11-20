package main

import "fmt"

func main() {
	gases := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
	}
	solids := map[string]map[string]string{
		"Fe": {
			"name":  "Iron",
			"state": "solid",
		},
		"Au": {
			"name":  "Gold",
			"state": "solid",
		},
	}

	var symbols string
	fmt.Println("Enter element symbols separated by spaces:")
	fmt.Scanln(&symbols)
	if el, ok := gases[symbols]; ok {
		fmt.Println(el["name"], el["state"])
	} else if !ok {
		fmt.Println("Element not found in gases, checking solids...")
		fmt.Scanln(&symbols)
		if el, ok := solids[symbols]; ok {
			fmt.Println(el["name"], el["state"])
		}
	} else {
		fmt.Println("Element not found")
	}
}
