package integration

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron"
)

const (
	CRON = "CRON"
	FILE_DROP = "FILE_DROP"
)

var DB *sql.DB

func Start(managers []Manager) {
	wg := sync.WaitGroup{}
	cron := cron.New()
	cron.Start()

	for _, manager := range managers {
		var err error
		DB, err = sql.Open("sqlite3", "./integration.db")
		if err != nil {
			log.Fatalln("Cannot open the database 'integration.db'. Error", err.Error())
		}

		//Get manager trigger information
		result := DB.QueryRow(`
			select t.type, t.param
			from managers m
			inner join triggers t on t.manager = m.name
			where m.active = 1 and m.name = ?
		`, manager.Name)

		if result.Err() != nil {
			log.Fatalln("Error while quering for the manager ", manager.Name, ". Error: ", result.Err().Error())
		}

		triggerType := ""
		param := ""
		err = result.Scan(&triggerType, &param)
		if err != nil {
			log.Fatalln("There is no manager with the name ", manager.Name, " in the database")
		}

		switch triggerType {
		case CRON:
			cron.AddFunc(param, func() {
				defer func() {
					if r := recover(); r != nil {
						log.Println("Recovered from manager", manager.Name)
					}
				}()
				manager.run()
			})

		case FILE_DROP:
		}
	}

	wg.Add(1)
	wg.Wait()
}
