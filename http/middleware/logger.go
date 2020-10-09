package middleware

import (
	"io"
	"log"
	"net/http"
	"time"

	"github.com/codegangsta/negroni"
)

// Logger is a middleware handler that logs the request as it goes in and the response as it goes out.
type Logger struct {
	// Logger inherits from log.Logger used to log messages with the Logger middleware
	*log.Logger
}

// NewLogger returns a new Logger instance
func NewLogger(out io.Writer) *Logger {
	return &Logger{log.New(out, "", 0)}
}

func (l *Logger) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(rw, r)
	res := rw.(negroni.ResponseWriter)
	format := "[status:%d][use:%v][method:%s][uri:%s]"
	l.Printf(format, res.Status(), time.Since(start), r.Method, r.URL.RequestURI())
}
