package period

import (
	"github.com/nuriofernandez/WebArgus/memory/storage"
	"github.com/nuriofernandez/WebArgus/orders"
)

func GetAllCheckeable() []orders.Order {
	var result = make([]orders.Order, 0)

	for _, order := range storage.GetAll() {
		if ShouldBeChecked(order.Id) {
			result = append(result, order)
		}
	}

	return result
}
