package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	//variable for bank balance
	var bankBalance int
	var balance sync.Mutex

	// print out starting values
	fmt.Printf("inicial account balance: $%d.00", bankBalance)
	fmt.Println()

	// define weekly revenue
	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))
	// loop through 52 weeks and print out how much is med; keep a running total
	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()
			for week := 1; week <= 52; week++ {
				balance.Lock()

				temp := bankBalance
				temp += income.Amount
				bankBalance = temp

				balance.Unlock()
				fmt.Printf("On week %d, you earned $%d.00 from %s \n", week, income.Amount, income.Source)
			}
		}(i, income)
	}

	wg.Wait()
	//print out final balance

	fmt.Printf("Final bank balance: $%d.00", bankBalance)
	fmt.Println()
}

/*
func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()
	//thread safe operation
	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(2)

	go updateMessage("Hello, Universe", &mutex)
	go updateMessage("Hello, Cosmos!", &mutex)

	wg.Wait()

	fmt.Println(msg)
}

*/
