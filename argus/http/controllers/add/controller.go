package add

import (
	"encoding/json"
	"github.com/nuriofernandez/WebArgus/memory/storage"
	"github.com/nuriofernandez/WebArgus/orders"
	"io"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	var request Request
	json.Unmarshal(reqBody, &request)

	if !isValid(request) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Error: "Invalid request"})
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)

	newOrder := orders.Create(
		request.Url,
		request.CheckType,
		request.Period,
		request.Notify.Phone,
		request.Notify.Title,
		request.Notify.Message,
	)
	// Store it
	storage.Put(newOrder)

	json.NewEncoder(w).Encode(Response{
		Message: "Successfully created",
		UUID:    newOrder.Id,
	})
}

func isValid(request Request) bool {
	if request.Url == "" {
		return false
	}
	if request.Notify.Phone == "" {
		return false
	}
	if request.Notify.Title == "" {
		return false
	}
	if request.Notify.Message == "" {
		return false
	}
	// TODO: Check type and period (not implemented yet)
	return true
}
