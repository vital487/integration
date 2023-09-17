package managers

import (
	"github.com/vital487/integration/integration"	
	"github.com/vital487/integration/workers"
)


var SumManager integration.Manager = integration.Manager{
	Name: "SumManager",
	BeforeWorker: func(data interface{}) bool { return true },
	AfterWorker: func(data interface{}) bool { return true },
	Workers: []integration.Worker{
		workers.SumWorker,
		workers.WriteWorker,
	},
}