package integration

type WorkerResult struct {
	Data      interface{}
	Result    bool
	Message   string
	KeepGoing bool
}