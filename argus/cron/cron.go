package cron

import (
	"fmt"
	"github.com/nuriofernandez/WebArgus/memory/storage"
	"time"
)

func Start() {
	fmt.Println("[Cron] Preparing orders checker...")

	for {
		orders := storage.GetAll()
		if len(orders) == 0 {
			fmt.Println("[Cron] There is no orders available")
			return
		}

		for _, order := range orders {
			Iterate(order)
		}

		fmt.Println("[Cron] Operation completed.")

		time.Sleep(time.Second * 5)
		fmt.Println("[Cron] Rechecking in 1 hour...")

		// wait 1h
		time.Sleep(time.Hour)
	}
}
