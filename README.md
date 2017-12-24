# call911

call911 is a Go middleware to display debug information when an application is in a panic situation
or needs to handle an error.

This middleware contains two features:

* An enhanced console output of the error stacktrace using [panicparse](https://github.com/maruel/panicparse)
* An HTML debug error page which contains stacktrace and various information about request, environment, etc.
