package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	//whatever come after this, don't execute it until the current function exists.
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	//waitgroup
	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
	}

	//wg.Add(len(words))

	for i, x := range words {
		wg.Add(1)
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait()

	wg.Add(1)

	printSomething("This is the second thing", &wg)
}
