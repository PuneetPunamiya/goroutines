package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_PrintSomething(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)

	go printSomething("data", &wg)

	wg.Wait()

	_ = w.Close()

	res, _ := io.ReadAll(r)
	output := string(res)

	if !strings.Contains(output, "data") {
		t.Errorf("Bhai kuch toh lafda hai")
	}
}
