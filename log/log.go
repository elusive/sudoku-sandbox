package log

import (
    "fmt"
    "io"
    "os"
)

// use to catch output
var w io.Writer

// Verbose enables debug.
var Verbose = false

// Width of column.
var Width = "14"

// Set alternate output
func Init(out io.Writer) {
    w = out
}

// Info.
func Info(key string, format string, args ...interface{}) {
    p := fmt.Sprintf("\033[36m%"+Width+"s \033[90m:\033[0m ", key)
    fmt.Fprintf(w, p+format+"\n", args...)
}

// Debug.
func Debug(key string, format string, args ...interface{}) {
    if Verbose {
        p := fmt.Sprintf("\033[96m%"+Width+"s \033[90m:\033[0m ", key)
        fmt.Fprintf(w, p+format+"\n", args...)
    }
}

// Warning.
func Warn(format string, args ...interface{}) {
    p := fmt.Sprintf("\033[33m%"+Width+"s \033[90m:\033[0m ", "warn")
    fmt.Fprintf(w, p+format+"\n", args...)
}

// Error.
func Error(err error) {
    p := fmt.Sprintf("\033[31m%"+Width+"s \033[90m:\033[0m ", "error")
    fmt.Fprintf(os.Stderr, p+"%s\n", err.Error())
}
