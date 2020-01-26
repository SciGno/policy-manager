package logger

import (
	"fmt"
	"log"
	"os"
)

var info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lmicroseconds)
var warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lmicroseconds)
var debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
var errors = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lmicroseconds)
var fatal = log.New(os.Stdout, "ERROR: ", log.Lshortfile|log.LstdFlags|log.Ldate|log.Ltime|log.Lmicroseconds)
var audit = log.New(os.Stdout, "AUDIT: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

// Infof formated function
func Infof(format string, v ...interface{}) {
	info.Printf(fmt.Sprintf(format, v...))
}

// Info function
func Info(v ...interface{}) {
	info.Print(fmt.Sprint(v...))
}

// Warningf function
func Warningf(format string, v ...interface{}) {
	warning.Printf(fmt.Sprintf(format, v...))
}

// Warning function
func Warning(v ...interface{}) {
	warning.Print(fmt.Sprint(v...))
}

// Debugf formated function
func Debugf(format string, v ...interface{}) {
	debug.Printf(fmt.Sprintf(format, v...))
}

// Debug function
func Debug(v ...interface{}) {
	debug.Print(fmt.Sprint(v...))
}

// Auditf formated function
func Auditf(format string, v ...interface{}) {
	audit.Printf(fmt.Sprintf(format, v...))
}

// Audit function
func Audit(v ...interface{}) {
	audit.Print(fmt.Sprint(v...))
}

// Errorf formated function
func Errorf(format string, v ...interface{}) {
	errors.Printf(fmt.Sprintf(format, v...))
}

// Error function
func Error(v ...interface{}) {
	errors.Print(fmt.Sprint(v...))
}

// Fatalf formated function
func Fatalf(format string, v ...interface{}) {
	fatal.Printf(fmt.Sprintf(format, v...))
}

// Fatal function
func Fatal(v ...interface{}) {
	fatal.Print(fmt.Sprint(v...))
}
