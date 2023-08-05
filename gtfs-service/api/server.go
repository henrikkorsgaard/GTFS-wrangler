package api

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Need a proper abstraction for this
// Do I run serve in a go routine
// or do I run it in a go routine within the serve function?
func Serve() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	http.ListenAndServe(":3333", r)
}