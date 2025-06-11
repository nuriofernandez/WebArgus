package list

import (
	"encoding/json"
	"github.com/nuriofernandez/WebArgus/memory/storage"
	"github.com/nuriofernandez/WebArgus/orders"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	allOrders := storage.GetAll()
	response := convert(allOrders)

	json.NewEncoder(w).Encode(response)
}

func convert(orders []orders.Order) []OrderResponse {
	var result []OrderResponse
	for _, order := range orders {
		result = append(result, OrderResponse{
			Id:        order.Id,
			Url:       order.Url,
			CheckType: order.CheckType,
			Notify: NotificationRuleResponse{
				Phone:   order.Notify.Phone,
				Title:   order.Notify.Title,
				Message: order.Notify.Message,
			},
			Period: order.Period,
		})
	}
	return result
}
