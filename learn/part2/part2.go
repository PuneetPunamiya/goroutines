package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	url := []string{
		"https://google.com",
		"https://github.com",
		"https://api.hub.tekton.dev",
		"https://fb.com",
	}

	for _, u := range url {
		wg.Add(1)

		go func() {
			defer wg.Done()
			resp, err := http.Get(u)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Printf("Status of url %s is := %s \n", u, resp.Status)
		}()
	}

	wg.Wait()
}
