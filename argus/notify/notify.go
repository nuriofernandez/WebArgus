package notify

import (
	"fmt"
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

	fmt.Println("[Notify] Sending notification...")
	sendSms(info.Phone, info.Title, info.Message)
}
