package workers

import (
	"log"
	"github.com/vital487/integration/integration"
)

var WriteWorker integration.Worker = integration.Worker{
	Name: "WriteWorker",
	Run: func(params map[string]interface{}, data interface{}) integration.WorkerResult{
		a  := data.(float64)
		result := integration.WorkerResult{}

		log.Println(a)

		result.KeepGoing = true
		result.Result = true
		return result
	},
}

