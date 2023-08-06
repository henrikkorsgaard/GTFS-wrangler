package api

import (
	"net/http"
	"github.com/go-chi/chi"
)

// Need a proper abstraction for this
// Do I run serve in a go routine
// or do I run it in a go routine within the serve function?
func StartServer() {
	r := chi.NewRouter()

	r.Get("/api/rest/stops", StopHandler)
	
	http.ListenAndServe(":8080", r)
}