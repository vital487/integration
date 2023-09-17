package main

import (
	"github.com/vital487/integration/integration"
	"github.com/vital487/integration/managers"
)

func main() {
	managers := []integration.Manager{
		//managers.SumManager,
		managers.DbManager,
	}

	integration.Start(managers)
}
