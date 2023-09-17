package integration

type Worker struct {
	Name string
	Run  func(params map[string]interface{}, resultData interface{}) WorkerResult
}