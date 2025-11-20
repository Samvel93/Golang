package main

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 8, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	min := x[0]
	minIdx := 1
	for i, v := range x {
		if v < min {
			min = v
			minIdx = i
		}
	}
	println("min:", min, "at index", minIdx)
}
