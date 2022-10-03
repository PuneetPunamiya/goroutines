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
	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("Bank balance of your account is %d", bankBalance)
	fmt.Println()

	var incomes = []Income{
		{Source: "Main job", Amount: 100},
		{Source: "Part time job", Amount: 10},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {

		go func(i int, income Income) {
			defer wg.Done()
			// Calculate for 52 weeks
			for week := 1; week <= 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week %d, you earned %d from income source %s\n", week, income.Amount, income.Source)
			}

		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Final bank balance is %d", bankBalance)
	fmt.Println()
}
