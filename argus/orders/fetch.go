package orders

import (
	"fmt"
	"os"
)

func NextOrder() (Order, error) {
	fmt.Println("[Order Handler] Preparing next order...")
	return Order{
		Url:       "https://itpec.org/statsandresults/all-passers-information/Philippines/2025S_FE.pdf",
		CheckType: "Status:200",
		Notify: NotificationRule{
			Phone:   "+34" + os.Getenv("TEST_ORDER_PHONE"),
			Title:   "ITPEC",
			Message: "ITPEC exam results are now available!",
		},
		Period: "1h",
	}, nil
}
