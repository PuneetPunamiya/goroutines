package main

import (
	"sync"
	"testing"
)

func Test_UpdateMessage(t *testing.T) {
	msg = "Hello, bidu"
	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Chaloooo bidu bye", &mutex)
	go updateMessage("Chaloooo bidu bye", &mutex)
	wg.Wait()
}
