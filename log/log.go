package log

import (
	"fmt"
	"log"
	"os"
)

var (
	info *log.Logger
	warning *log.Logger
	error *log.Logger
)

func init() {
	flags := log.Ldate | log.Ltime | log.Llongfile
	info = log.New(os.Stdout, "[INFO]", flags)
	warning = log.New(os.Stderr, "[WARN]", flags)
	error = log.New(os.Stderr, "[ERROR]", flags)
}

func Info(format string, v... interface{}) {
	info.Output(3, fmt.Sprintf(format, v...))
}

func Warning(format string, v... interface{}) {
	warning.Output(3, fmt.Sprintf(format, v...))
}

func Error(format string, v... interface{}) {
	error.Output(3, fmt.Sprintf(format, v...))
}

