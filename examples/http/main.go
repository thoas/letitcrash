package main

import (
	"fmt"
	"net/http"

	"github.com/thoas/letitgo"
)

func main() {
	h := letitgo.HandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
		return fmt.Errorf("looks like we have a panic situation")
	})

	http.Handle("/panic", letitgo.Handler(h, letitgo.WithVerbose(false)))
	http.ListenAndServe(":8080", nil)
}
