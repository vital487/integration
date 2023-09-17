package managers

import (
	"github.com/vital487/integration/integration"	
	"github.com/vital487/integration/workers"
)


var DbManager integration.Manager = integration.Manager{
	Name: "DbManager",
	BeforeWorker: func (data interface{}) bool { return true},
	AfterWorker: func (data interface{}) bool { return true},
	Workers: []integration.Worker{
		workers.GetDataWorker,
	},
}