package notify

import (
	"github.com/nuriofernandez/WebArgus/memory/archive"
	"github.com/nuriofernandez/WebArgus/memory/storage"
	"github.com/nuriofernandez/WebArgus/orders"
)

func Notify(order orders.Order) {
	// Archive it
	archive.Put(order)

	// Delete it
	storage.Delete(order.Id)

	// Send notification
	info := order.Notify
	sendSms(info.Phone, info.Title, info.Message)
}
