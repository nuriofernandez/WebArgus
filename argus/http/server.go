package http

import (
	"fmt"
	"github.com/nuriofernandez/WebArgus/http/controllers/defaultpath"
	"github.com/nuriofernandez/WebArgus/http/controllers/list"
	"net/http"
)

func Start() {
	// Prepare default path
	http.Handle("/", http.HandlerFunc(defaultpath.DefaultController))
	http.Handle("/orders", http.HandlerFunc(list.Controller))

	// Start http server
	fmt.Println("[HTTP] Listening on :80")
	http.ListenAndServe(":80", nil)
}
