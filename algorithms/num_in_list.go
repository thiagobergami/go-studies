package main

import "fmt"

func NumInList(list []int, num int) bool {
	for x := range list {
		if list[x] == num {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(NumInList([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(NumInList([]int{3, 3, 3, 3, 3}, 5))
	fmt.Println(NumInList([]int{3, 3, 3, 5, 3}, 5))
	fmt.Println(NumInList([]int{4, 2, 22, -10, 8}, -10))
	fmt.Println(NumInList(nil, 5))
	fmt.Println(NumInList([]int{}, 5))
}
