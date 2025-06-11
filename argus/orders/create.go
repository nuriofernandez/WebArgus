package orders

import "github.com/google/uuid"

func Create(url, checkType, period, phone, title, message string) Order {
	return Order{
		Id:        uuid.New().String(),
		Url:       url,
		CheckType: checkType,
		Notify: NotificationRule{
			Phone:   phone,
			Title:   title,
			Message: message,
		},
		Period: period,
	}
}
