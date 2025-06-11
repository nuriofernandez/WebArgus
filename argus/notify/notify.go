package notify

import (
	"github.com/nuriofernandez/WebArgus/memory"
	"github.com/nuriofernandez/WebArgus/orders"
)

func Notify(order orders.Order) {
	isChecked, err := memory.Check(order.Url)

	// In case is already sent, stop.
	if err != nil || !isChecked {
		return
	}

	info := order.Notify
	err = memory.Remember(order.Url)
	if err != nil {
		return
	}
	sendSms(info.Phone, info.Title, info.Message)
}
