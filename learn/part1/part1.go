package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var a string
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		fmt.Println("I'm sleeping for 10 sec")
		time.Sleep(10 * time.Second)

		a = "hello world"
		fmt.Println("Done sleeping")

		wg.Done()
	}()

	a = "before hello"
	fmt.Println(a)

	wg.Wait()
	fmt.Println(a)
}
