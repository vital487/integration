package integration

import (
	"errors"
	"log"
	"strconv"
)

type Manager struct {
	Name         string
	Workers      []Worker
	BeforeWorker func(data interface{}) bool
	AfterWorker  func(data interface{}) bool
}

func (manager Manager) run() {
	log.Println("Starting Manager", manager.Name)

	var data interface{}

	for _, worker := range manager.Workers {
		log.Println("Starting Worker", worker.Name)

		//Run BeforeWorker
		beforeResult := manager.BeforeWorker(data)
		if !beforeResult {
			log.Println("Result of BeforeWorker was false, before the worker", worker.Name)
			break
		}

		//Get worker parameters
		params, err := getWorkerParams(manager.Name, worker.Name)
		if err != nil {
			log.Println(err.Error())
			break
		}

		//Run worker
		result := worker.Run(params, data)
		data = result.Data

		//Log the result and the message
		if !result.Result {
			log.Println("Result of Worker", worker.Name+": Failure")
		}

		if result.Message != "" {
			log.Println("Message of Worker", worker.Name+": ", result.Message)
		}

		if !result.KeepGoing {
			log.Println("Worker", worker.Name, "stopped the execution of the next workers")
			log.Println("Finishing Worker", worker.Name)
			break
		}

		//Run AfterWorker
		afterResult := manager.AfterWorker(data)
		if !afterResult {
			log.Println("Result of AfterWorker was false, after the worker", worker.Name)
			break
		}

		// log.Println("Finishing Worker", worker.Name)
	}

	// log.Println("Finishing Manager", manager.Name)
}

func getWorkerParams(manager string, worker string) (map[string]interface{}, error) {
	params := make(map[string]interface{})

	//Get worker parameters
	stmt, err := DB.Prepare(`
		select p.type, p.name, p.value
		from workers w
		inner join managers m on m.name = w.manager
		inner join parameters p on p.worker = w.id
		where m.name = ? and w.name = ?
	`)
	if err != nil {
		return nil, errors.New("Failed to check the parameters of the worker " + worker + ". Error: " + err.Error())
	}

	result, err := stmt.Query(manager, worker)
	if err != nil {
		return nil, errors.New("Getting parameters for worker " + worker + " failed. Error: " + err.Error())
	}

	//Add parameters to params with the respective type
	for result.Next() {
		var paramType string
		var name string
		var value string

		result.Scan(&paramType, &name, &value)

		switch paramType {
		case "int":
			params[name], err = strconv.Atoi(value)
			if err != nil {
				return nil, errors.New("Failed to convert param " + name + " from worker " + worker + " to int. Error: " + err.Error())
			}
		case "float":
			params[name], err = strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, errors.New("Failed to convert param " + name + " from worker " + worker + " to float64. Error: " + err.Error())
			}
		case "string":
			fallthrough
		default:
			params[name] = value
		}
	}

	return params, nil
}
