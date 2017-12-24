package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/pkg/errors"
	"github.com/thoas/call911"
)

func main() {
	verboseFlag := false

	log.SetFlags(log.Lmicroseconds)
	if !verboseFlag {
		log.SetOutput(ioutil.Discard)
	}

	defer func() {
		if r := recover(); r != nil {
			call911.HandleError(r)
		}
	}()

	call911.HandleError(errors.WithStack(fmt.Errorf("generated error")))

	panic(errors.WithStack(fmt.Errorf("it's called panic")))
}
