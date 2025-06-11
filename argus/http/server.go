package http

import (
	"fmt"
	"github.com/nuriofernandez/WebArgus/http/controllers/defaultpath"
	"net/http"
)

func Start() {
	// Prepare default path
	http.Handle("/", http.HandlerFunc(defaultpath.DefaultController))

	// Start http server
	fmt.Println("[HTTP] Listening on :80")
	http.ListenAndServe(":80", nil)
}
