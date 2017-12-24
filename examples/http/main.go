package main

import (
	"fmt"
	"net/http"

	"github.com/thoas/call911"
)

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic(fmt.Errorf("looks like we have a panic situation"))
	})

	http.Handle("/panic", call911.Handler(h, call911.WithVerbose(false)))
	http.ListenAndServe(":8080", nil)
}
