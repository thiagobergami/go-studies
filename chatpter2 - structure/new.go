package main

import "fmt"

func newInt() *int {
	return new(int)
}

func main() {
	p := new(int)   //p, do tipo *int, aponta para uma variável int sem nome
	fmt.Println(*p) // 0
	*p = 2          // define o inte sem nome com o valor 2
	fmt.Println(*p) // 2

	// cada variável criada com new devolve uma variável distinta com um endereço único

	v := newInt()
	q := newInt()
	fmt.Println(v)
	fmt.Println(q)

}
