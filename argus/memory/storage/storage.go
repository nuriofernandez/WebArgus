package storage

import "github.com/nuriofernandez/WebArgus/orders"

var storage = make(map[string]orders.Order)

func Put(order orders.Order) {
	storage[order.Id] = order
}

func Get(id string) orders.Order {
	return storage[id]
}

func Exist(id string) bool {
	_, exists := storage[id]
	return exists
}

func GetAll() []orders.Order {
	var result []orders.Order
	for _, order := range storage {
		result = append(result, order)
	}
	return result
}
