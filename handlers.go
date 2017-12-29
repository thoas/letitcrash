package letitgo

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// ErrorHandler is an handler to render the error.
type ErrorHandler func(err interface{}, w http.ResponseWriter, r *http.Request)

// DefaultErrorHandler is an handler provided by letitgo to render the error.
func DefaultErrorHandler(err interface{}, w http.ResponseWriter, r *http.Request) {
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
