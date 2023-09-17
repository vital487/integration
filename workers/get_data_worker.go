package workers

import (
	"fmt"

	"github.com/vital487/integration/integration")


var GetDataWorker integration.Worker = integration.Worker{
	Name: "GetDataWorker",
	Run: func(params map[string]interface{}, resultData interface{}) integration.WorkerResult {
		result := integration.WorkerResult{}

		connection, err := integration.GetConnectionFromCompanyName(params["targetCompany"].(string))
		if err != nil {
			result.KeepGoing = false
			// result.Message = "There is no company configured with the name " + params["targetCompany"].(string)
			result.Message = err.Error()
			return result
		}

		db, err := connection.GetDb()
		if err != nil {
			result.KeepGoing = false
			// result.Message = "Error while connecting to the database of the company " + params["targetCompany"].(string)
			result.Message = err.Error()
			return result
		}

		//Get data
		data, err := db.Query(`
			select CardName
			from OCRD
		`)
		if err != nil {
			result.KeepGoing = false
			result.Message = err.Error()
			return result
		}

		users := []string{}
		for data.Next() {
			var name string
			data.Scan(&name)
			users = append(users, name)
		}

		fmt.Println(users)

		result.Data = users
		result.KeepGoing = true
		return result
	},
}