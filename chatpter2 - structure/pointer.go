package main

import "fmt"

func f() *int {
	v := 1
	return &v
}

func main() {
	var p = f()
	x := 2
	v := &x

	fmt.Println(*p) // 1 => pointer value
	fmt.Println(v)  // pointer address
	fmt.Println(&x) // pointer address

}
