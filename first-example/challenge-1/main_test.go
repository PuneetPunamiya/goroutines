package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_UpdateMessage(t *testing.T) {
	wg.Add(1)
	go updateMessage("Works")
	wg.Wait()

	if msg != "Works" {
		t.Errorf("It did not works properly")
	}
}

func Test_PrintMessage(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Go routine rocks"
	printMessage()

	_ = w.Close()

	output, _ := io.ReadAll(r)
	res := string(output)

	if !strings.Contains(res, "Go routine rocks") {
		t.Error("Expected to find `Go routine rocks`, but it is not there")
	}
}
