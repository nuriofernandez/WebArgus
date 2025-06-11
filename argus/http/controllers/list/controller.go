package list

import (
	"encoding/json"
	"net/http"
	"os"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(OrderResponse{
		Id:        "ABC",
		Url:       "https://itpec.org/statsandresults/all-passers-information/Philippines/2025S_FE.pdf",
		CheckType: "Status:200",
		Notify: NotificationRuleResponse{
			Phone:   "+34" + os.Getenv("TEST_ORDER_PHONE"),
			Title:   "ITPEC",
			Message: "ITPEC exam results are now available!",
		},
		Period: "1h",
	})
}
