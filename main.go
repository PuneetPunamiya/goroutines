package main

import (
	"encoding/json"
	"log"
)

func main() {
	noOfJobs := 3000
	go allocateJobs(noOfJobs)

	done := make(chan bool)
	go getResults(done)

	noOfWorkers := 100
	createWorkerPool(noOfWorkers)

	<-done

	// convert result collection to JSON
	data, err := json.MarshalIndent(resultCollection, "", "    ")
	if err != nil {
		log.Fatal("json err: ", err)
	}

	if err := writeToFile(data); err != nil {
		log.Fatal("writeToFile err: ", err)
	}

}
