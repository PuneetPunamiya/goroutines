package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	leftFork  int
	rightFork int
}

var philosophers = []Philosopher{
	{name: "Karan", leftFork: 4, rightFork: 0},
	{name: "Adish", leftFork: 0, rightFork: 1},
	{name: "Bineet", leftFork: 1, rightFork: 2},
	{name: "Rohan", leftFork: 2, rightFork: 3},
	{name: "Puneet", leftFork: 3, rightFork: 4},
}

// Define a few variables.
var hunger = 3                  // how many times a philosopher eats
var eatTime = 1 * time.Second   // how long it takes to eatTime
var thinkTime = 3 * time.Second // how long a philosopher thinks
var sleepTime = 1 * time.Second // how long to wait when printing things out

func main() {
	// Print the welcome message
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("---------------------------")
	fmt.Println("The table is empty")

	// start the meal
	dine()
}

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// Forks is a map of all 5 forks
	var forks = make(map[int]*sync.Mutex)
	for i := range philosophers {
		forks[i] = &sync.Mutex{}
	}

	// start the meal
	for i := range philosophers {
		go diningProblem(philosophers[i], forks, wg, seated)
	}
	wg.Wait()
}

func diningProblem(philosopher Philosopher, forks map[int]*sync.Mutex, wg, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat the philosopher at the table
}
