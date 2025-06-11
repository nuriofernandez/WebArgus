package cron

import (
	"fmt"
	"github.com/nuriofernandez/WebArgus/cron/period"
	"time"
)

func Start() {
	fmt.Println("[Cron] Preparing orders checker...")

	for {
		// Wait 5 min to recheck
		time.Sleep(time.Minute * 5)

		orders := period.GetAllCheckeable()
		if len(orders) == 0 {
			fmt.Println("[Cron] There is no orders available")
			continue
		}

		for _, order := range orders {
			Iterate(order)
		}

		fmt.Println("[Cron] Operation completed.")

		time.Sleep(time.Second * 5)
		fmt.Println("[Cron] Rechecking in 5 minutes...")
	}
}
