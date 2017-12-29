package main

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"github.com/thoas/letitgo"
)

func main() {
	h1 := letitgo.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return fmt.Errorf("looks like we have a panic situation")
	})

	h2 := letitgo.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return errors.WithStack(fmt.Errorf("looks like we have a panic situation"))
	})

	http.Handle("/panic", letitgo.Handler(h1, letitgo.WithVerbose(false)))
	http.Handle("/panicstack", letitgo.Handler(h2, letitgo.WithVerbose(false)))
	http.ListenAndServe(":8080", nil)
}
