package main

import "fmt"

func Sum(numbers []int) int {
	var sum int
	for _, val := range numbers {
		sum += val
	}

	return sum
}

func main() {
	fmt.Println(Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(Sum([]int{3, 3, 3, 3, 3}))
	fmt.Println(Sum([]int{3, 5, 3, 5, 3}))
	fmt.Println(Sum(nil))
	fmt.Println(Sum([]int{}))
}
