package transitworker

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetupRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/v1/:mount/encrypt/:name", Hello)

	return router
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s:%s!\n", ps.ByName("name"), ps.ByName("mount"))
}
