package call911

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type ErrorHandler func(err interface{}, w http.ResponseWriter, r *http.Request)

func defaultErrorHandler(err interface{}, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)

	queryString := make(map[string]string)

	for key, value := range r.URL.Query() {
		queryString[key] = fmt.Sprintf("%+v", value)
	}

	t := template.Must(template.New("debug").Parse(debugTpl))

	headers := make(map[string]string)

	for key, value := range r.Header {
		if key == "Cookie" {
			continue
		}

		headers[key] = fmt.Sprintf("%+v", value)
	}

	environments := make(map[string]string)

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		environments[pair[0]] = pair[1]
	}

	t.Execute(w, map[string]interface{}{
		"queryString":       queryString,
		"request":           r,
		"error":             err,
		"cookies":           r.Cookies(),
		"version":           runtime.Version(),
		"serverInformation": getServerInformation(),
		"now":               time.Now(),
		"errorType":         reflect.TypeOf(err).String(),
		"headers":           headers,
		"environments":      environments,
	})
}

var defaultOptions = newOptions(
	WithVerbose(false),
	WithConsole(true),
	WithErrorHandler(defaultErrorHandler),
)

// SetDefaultOptions sets default options.
func SetDefaultOptions(opts ...Option) {
	defaultOptions = newOptions(opts...)
}

// HandleRecover handles an error from a recover situation.
func HandleRecover(w http.ResponseWriter, r *http.Request, opts ...Option) {
	if err := recover(); err != nil {
		HandleError(err, w, r, opts...)
	}
}

// HandleError handles an error.
func HandleError(err interface{}, w http.ResponseWriter, r *http.Request, opts ...Option) {
	opt := defaultOptions.Merge(newOptions(opts...))

	if opt.Verbose != nil && !*opt.Verbose {
		log.SetOutput(ioutil.Discard)
	}

	if opt.Console != nil && *opt.Console {
		PrintError(err)
	}

	opt.ErrorHandler(err, w, r)
}

// Handler is a MiddlewareFunc which implements the Middleware interface.
func Handler(h http.Handler, opts ...Option) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer HandleRecover(w, r, opts...)

		h.ServeHTTP(w, r)
	})
}

// HandlerFunc is a http.HandlerFunc which can return an error.
type HandlerFunc func(http.ResponseWriter, *http.Request) error

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		HandleError(err, w, r)
	}
}

// ServeHTTP is a Negroni compatible interface.
func ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc, opts ...Option) {
	defer HandleRecover(w, r, opts...)

	next(w, r)
}
