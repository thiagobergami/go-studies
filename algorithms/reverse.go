package main

import "fmt"

func Reverse(word string) string {
	var res string
	for i := len(word) - 1; i >= 0; i-- {
		res = res + string(word[i])
	}

	return res
}

func main() {
	fmt.Println(Reverse("cat"))
	fmt.Println(Reverse("alphabet"))
}
