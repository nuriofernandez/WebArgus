package defaultpath

import (
	"encoding/json"
	"net/http"
)

func DefaultController(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(DefaultResponse{Error: "Resource not specified"})
}
