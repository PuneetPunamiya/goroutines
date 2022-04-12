package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func allocateJobs(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		jobs <- Job{i + 1}
	}
	close(jobs)
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		result, err := fetch(job.number)
		if err != nil {
			log.Printf("error in fetching: %v\n", err)
		}
		results <- *result
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func getResults(done chan bool) {
	for res := range results {
		if res.Num != 0 {
			resultCollection = append(resultCollection, res)
		}
	}
	done <- true
}

// Client -> http Request -> Check the response -> get the data

func fetch(id int) (*Result, error) {
	client := http.Client{
		Timeout: 5 * time.Minute,
	}

	url := fmt.Sprintf("%s/%d/info.0.json", URL, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var data Result
	if resp.StatusCode != http.StatusOK {
		data = Result{}
	} else {
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return nil, err
		}
	}

	resp.Body.Close()

	return &data, nil
}

func writeToFile(data []byte) error {
	f, err := os.Create("xkcd.json")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
