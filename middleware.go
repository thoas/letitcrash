package call911

import (
	"io/ioutil"
	"log"
	"net/http"
)

// HandleRecover handles an error from a recover situation.
func HandleRecover(w http.ResponseWriter, r *http.Request, opts ...Option) {
	opt := newOptions(opts...)

	log.SetFlags(log.Lmicroseconds)
	if !opt.Verbose {
		log.SetOutput(ioutil.Discard)
	}

	if err := recover(); err != nil {
		HandleError(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`Internal server error`))
	}
}

// Handler is a MiddlewareFunc which implements the Middleware interface.
func Handler(h http.Handler, opts ...Option) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer HandleRecover(w, r, opts...)

		h.ServeHTTP(w, r)
	})
}

// ServeHTTP is a Negroni compatible interface.
func ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, opts ...Option) {
	defer HandleRecover(w, r, opts...)

	next(w, r)
}
