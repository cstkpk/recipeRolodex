package logger

import (
	"log"
	"os"
	"runtime"
	"strings"
)

var (
	// Info is for logging any useful information, i.e. success info
	Info *log.Logger
	// Debug is for logging while developing/debugging
	Debug *log.Logger
	// Error is for logging errors
	Error *log.Logger
)

// https://en.wikipedia.org/wiki/ANSI_escape_code#Colors
func init() {

	Info = log.New(os.Stdout,
		"\033[0;36mINFO: \033[0m",
		log.Ldate|log.Ltime|log.Lshortfile)

	Debug = log.New(os.Stdout,
		"\033[0;33mDEBUG: \033[0m",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(os.Stdout,
		"\033[0;31mERROR: \033[0m",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// GetCallInfo retrieves package and function name at the stack level where the log occurs
// TODO: Might eventually be useful to pass an int into the function and change
// runtime.Caller(1) to runtime.Caller(i) so info from any stack level could be logged, but
// this would require more validation
func GetCallInfo() string {
	pc, _, _, _ := runtime.Caller(1)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	funcName := parts[pl-1]

	filePathParts := strings.Split(parts[1], "/")
	pkg := filePathParts[len(filePathParts)-1]

	return pkg + "." + funcName + "():"
}
