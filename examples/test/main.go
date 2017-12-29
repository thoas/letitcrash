package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/thoas/letitgo"
)

func main() {
	verboseFlag := false

	log.SetFlags(log.Lmicroseconds)
	if !verboseFlag {
		log.SetOutput(ioutil.Discard)
	}

	defer func() {
		if r := recover(); r != nil {
			letitgo.PrintError(r, os.Stdout)
		}
	}()

	letitgo.PrintError(errors.WithStack(fmt.Errorf("generated error")), os.Stdout)

	panic(errors.WithStack(fmt.Errorf("it's called panic")))
}
