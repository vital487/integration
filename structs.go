package integration

import (
	"log"
)

const (
	CRON = iota
	FILE_DROP
)

type Trigger struct {
	trigger int8
	params  []string
}

type Runnable interface {
	run()
}

type WorkerResult struct {
	data      interface{}
	result    bool
	message   string
	keepGoing bool
}

type Worker struct {
	name string
	run  func(data interface{}) WorkerResult
}

type Manager struct {
	trigger Trigger
	name    string
	workers []Worker
}

func (manager Manager) run() {
	log.Println("Starting Manager ", manager.name)

	var data interface{}

	for _, worker := range manager.workers {
		log.Println("Starting Worker ", worker.name)

		//Run worker
		result := worker.run(data)
		data = result.data

		//Log the result and the message
		if result.result {
			log.Println("Result of Worker ", worker.name, ": Success")
		} else {
			log.Println("Result of Worker ", worker.name, ": Failure")
		}

		if result.message != "" {
			log.Println("Message of Worker ", worker.name, ": ", result.message)
		}

		if !result.keepGoing {
			log.Println("Worker ", worker.name, " stopped the execution of the next workers")
			log.Println("Finishing Worker ", worker.name)
			break
		}

		log.Println("Finishing Worker ", worker.name)
	}

	log.Println("Finishing Manager ", manager.name)
}
