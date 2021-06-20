package integration

import (
	"sync"
	"github.com/robfig/cron"
)

func Start(managers []Manager) {
	wg := sync.WaitGroup{}
	cron := cron.New()

	for _, manager := range managers {
		switch manager.trigger.trigger {
		case CRON:
			cron.AddFunc("30 * * * ?", func() {
				manager.run()
			})

		case FILE_DROP:
		}
	}

	wg.Add(1)
	wg.Wait()
}
