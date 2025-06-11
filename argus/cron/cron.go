package cron

import (
	"fmt"
	"github.com/nuriofernandez/WebArgus/checker"
	"github.com/nuriofernandez/WebArgus/notify"
	"github.com/nuriofernandez/WebArgus/orders"
)

func Start() {
	fmt.Println("[Cron] Preparing orders checker...")
	order, err := orders.NextOrder()
	if err != nil {
		fmt.Println("[Cron] There is no orders available")
		return
	}

	isOnline := checker.IsOnline(order.Url)
	if isOnline {
		fmt.Println("[Cron] Order passed check!")

		fmt.Println("[Cron] Preparing notification...")
		notify.Notify(order)
	}

	fmt.Println("[Cron] Operation completed.")
}
