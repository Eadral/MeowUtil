package log

import (
	"fmt"
	"log"
	"os"
)

var (
	infoLogger      *log.Logger
	warningLogger   *log.Logger
	errorLogger     *log.Logger
	disableOutput   bool
	conditionOutput bool
)

func init() {
	flags := log.Ldate | log.Ltime | log.Lshortfile
	infoLogger = log.New(os.Stdout, "[INFO]", flags)
	warningLogger = log.New(os.Stderr, "[WARN]", flags)
	errorLogger = log.New(os.Stderr, "[ERROR]", flags)

	if os.Getenv("MEOWLOG") == "" {
		disableOutput = true
	}
}

func ConditionOutput(yes bool) {
	conditionOutput = yes
}

func Info(format string, v ...interface{}) {
	if conditionOutput && !disableOutput {
		infoLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Warning(format string, v ...interface{}) {
	if conditionOutput && !disableOutput {
		warningLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func Error(format string, v ...interface{}) {
	if conditionOutput && !disableOutput {
		errorLogger.Output(2, fmt.Sprintf(format, v...))
	}
}

func ERR(format string, v ...interface{}) error {
	newErr := fmt.Errorf(format, v...)
	if conditionOutput && !disableOutput {
		errorLogger.Output(2, newErr.Error())
	}
	return newErr
}

func Fatal(format string, v ...interface{}) {
	if conditionOutput && !disableOutput {
		errorLogger.Output(2, fmt.Sprintf(format, v...))
	}
	os.Exit(1)
}
