package main

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/thoas/letitcrash"
)

func main() {
	h1 := letitcrash.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return fmt.Errorf("looks like we have a panic situation")
	})

	h2 := letitcrash.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return errors.WithStack(fmt.Errorf("looks like we have a panic situation"))
	})

	http.Handle("/panic", letitcrash.Handler(h1, letitcrash.WithVerbose(false)))
	http.Handle("/panicstack", letitcrash.Handler(h2, letitcrash.WithVerbose(false)))
	http.ListenAndServe(":8080", nil)
}
