package workers

import (
	"github.com/vital487/integration/integration"
)

var SumWorker integration.Worker = integration.Worker{
	Name: "SumWorker",
	Run: run,
}

func run(params map[string]interface{}, data interface{}) integration.WorkerResult{
	result := integration.WorkerResult{}

	num1 := params["Num1"].(float64)
	num2 := params["Num2"].(float64)

	result.Data = num1 * num2
	result.KeepGoing = true
	result.Result = true
	return result
}