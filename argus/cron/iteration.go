package cron

import (
	"fmt"
	"github.com/nuriofernandez/WebArgus/checker"
	"github.com/nuriofernandez/WebArgus/notify"
	"github.com/nuriofernandez/WebArgus/orders"
)

func Iterate(order orders.Order) {
	fmt.Printf("[Cron] (%s) Checking...\n", order.Id)

	isOnline := checker.IsOnline(order.Url)
	if isOnline {
		fmt.Printf("[Cron] (%s) Order passed check!\n", order.Id)

		fmt.Printf("[Cron] (%s) Preparing notification...\n", order.Id)
		notify.Notify(order)
	}

	fmt.Printf("[Cron] (%s) Iteration completed.\n", order.Id)
}
