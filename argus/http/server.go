package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nuriofernandez/WebArgus/http/controllers/add"
	"github.com/nuriofernandez/WebArgus/http/controllers/defaultpath"
	"github.com/nuriofernandez/WebArgus/http/controllers/list"
	"github.com/nuriofernandez/WebArgus/http/middlewares"
	"log"
	"net/http"
)

func Start() {
	// Create a new router
	apiRouter := mux.NewRouter().StrictSlash(true)

	// Register middlewares
	apiRouter.Use(
		middlewares.ContentTypeMiddleware,
		middlewares.LogRequest,
		middlewares.Authentication,
	)

	// Prepare default path
	apiRouter.HandleFunc("/", defaultpath.DefaultController)
	apiRouter.NotFoundHandler = http.HandlerFunc(defaultpath.DefaultController)

	// Prepare paths
	apiRouter.HandleFunc("/orders", list.Controller)
	apiRouter.HandleFunc("/orders/add", add.Controller)

	// Start http server
	fmt.Println("[HTTP] Listening on :80")
	log.Fatal(http.ListenAndServe(":80", apiRouter))
}
