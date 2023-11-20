package main

import "fmt"

func FizzBuzz(number int) {
	for i := 1; i <= number; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i)
		}
	}
}

func main() {
	FizzBuzz(1)
	FizzBuzz(5)
	FizzBuzz(15)
}
